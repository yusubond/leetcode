package template

import (
	"testing"
	"time"
)

// fakeClock 可控时钟，便于精确演示过期行为
type fakeClock struct{ t time.Time }

func newFakeClock() *fakeClock { return &fakeClock{t: time.UnixMilli(0)} }
func (f *fakeClock) now() time.Time { return f.t }
func (f *fakeClock) add(d time.Duration) { f.t = f.t.Add(d) }

func TestTTLCache_Basic(t *testing.T) {
	clk := newFakeClock()
	c := NewTTLCache(2, 3*time.Second)
	c.now = clk.now

	c.Put(1, 1) // expire @3
	c.Put(2, 2) // expire @4
	if got := c.Get(1); got != 1 {
		t.Fatalf("Get(1) = %d, want 1", got)
	}
	if c.Len() != 2 {
		t.Fatalf("Len = %d, want 2", c.Len())
	}
}

// 惰性删除：过期后 Get 返回 -1，并从缓存移除
func TestTTLCache_LazyExpiration(t *testing.T) {
	clk := newFakeClock()
	c := NewTTLCache(2, 3*time.Second)
	c.now = clk.now

	c.Put(1, 1)                  // expire @3
	clk.add(2 * time.Second)     // now=2，未过期
	if got := c.Get(1); got != 1 {
		t.Fatalf("before expiry Get(1) = %d, want 1", got)
	}
	clk.add(2 * time.Second) // now=4，已过期(3<=4)
	if got := c.Get(1); got != -1 {
		t.Fatalf("after expiry Get(1) = %d, want -1", got)
	}
	if c.Len() != 0 {
		t.Fatalf("after expiry Len = %d, want 0 (lazy deleted)", c.Len())
	}
}

// 超容量时优先淘汰过期项（这是 TTL 版相对 146 LRU 的关键差别）
// 1 在 t=3 过期、2 在 t=4 过期；t=3 插入 3 时，evictExpired 清掉过期的 1，2 被保留。
func TestTTLCache_EvictExpiredBeforeLRU(t *testing.T) {
	clk := newFakeClock()
	c := NewTTLCache(2, 3*time.Second)
	c.now = clk.now

	c.Put(1, 1)                  // expire @3
	clk.add(1 * time.Second)     // now=1
	c.Put(2, 2)                  // expire @4
	clk.add(2 * time.Second)     // now=3：1 已过期(3<=3)，2 未过期(4>3)
	c.Put(3, 3)                  // 超容量 → evictExpired 清掉 1，2 保留

	if got := c.Get(1); got != -1 {
		t.Fatalf("Get(1) = %d, want -1 (过期被主动清理)", got)
	}
	if got := c.Get(2); got != 2 {
		t.Fatalf("Get(2) = %d, want 2 (未过期应保留)", got)
	}
	if got := c.Get(3); got != 3 {
		t.Fatalf("Get(3) = %d, want 3", got)
	}
	if c.Len() != 2 {
		t.Fatalf("Len = %d, want 2", c.Len())
	}
}

// 更新已有 key 会重置 TTL 并 version++，旧堆条目变为脏数据
func TestTTLCache_UpdateResetsTTL(t *testing.T) {
	clk := newFakeClock()
	c := NewTTLCache(1, 3*time.Second)
	c.now = clk.now

	c.Put(1, 1)                  // expire @3, version 0
	clk.add(2 * time.Second)     // now=2
	c.Put(1, 11)                 // 命中未过期：val=11, expire 重置 @5, version 1
	clk.add(2 * time.Second)     // now=4：新 expireAt=5 仍未过期
	if got := c.Get(1); got != 11 {
		t.Fatalf("Get(1) = %d, want 11 (新值)", got)
	}
}

// 脏堆条目（已过期的旧 version）不应导致有效节点被误删
func TestTTLCache_StaleHeapEntrySkipped(t *testing.T) {
	clk := newFakeClock()
	c := NewTTLCache(2, 3*time.Second)
	c.now = clk.now

	c.Put(1, 1)                  // expire @3, v0
	clk.add(2 * time.Second)     // now=2
	c.Put(1, 11)                 // 命中未过期 → expire @5, v1；堆中 (3,v0) 变脏
	clk.add(2 * time.Second)     // now=4：1 真实过期时间 5 未到，但脏条目 (3,v0) 已"过期"
	c.Put(2, 2)                  // cache={1,2} 未超容量，不会自动触发清理

	c.evictExpired(clk.now())    // 手动清理：脏条目 (3,v0) 应被丢弃，有效节点 1 保留
	if c.Len() != 2 {
		t.Fatalf("Len = %d, want 2 (脏过期条目不应误删有效节点)", c.Len())
	}
	if got := c.Get(1); got != 11 {
		t.Fatalf("Get(1) = %d, want 11 (有效节点未被误删)", got)
	}
}
