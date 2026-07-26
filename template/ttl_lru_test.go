package template

import (
	"testing"
	"time"
)

func TestTTLLRUCache_Basic(t *testing.T) {
	clk := newFakeClock()
	c := NewTTLLRUCache(2, 3*time.Second)
	c.now = clk.now

	c.Put(1, 1) // expire @3
	c.Put(2, 2) // expire @4
	if got := c.Get(1); got != 1 {
		t.Fatalf("Get(1) = %d, want 1", got)
	}
	if c.Get(99) != -1 {
		t.Fatalf("Get(missing) should be -1")
	}
}

// 惰性删除 + 访问续期：命中会刷新过期时刻；真正过期后 Get 返回 -1 并移除
func TestTTLLRUCache_LazyExpirationWithRenew(t *testing.T) {
	clk := newFakeClock()
	c := NewTTLLRUCache(2, 3*time.Second)
	c.now = clk.now

	c.Put(1, 1)              // expire @3
	clk.add(2 * time.Second) // now=2，未过期
	if got := c.Get(1); got != 1 {
		t.Fatalf("Get(1) = %d, want 1", got)
	}
	// 访问续期：expire 由 3 刷新为 2+3=5；推进到 4（已过原始 expire 3），仍存活
	clk.add(2 * time.Second) // now=4
	if got := c.Get(1); got != 1 {
		t.Fatalf("Get(1) = %d, want 1 (续期让它活过原始过期时刻 3)", got)
	}
	// 上一次 Get(now=4) 再次续期到 4+3=7；停止访问，推过 7 即过期
	clk.add(4 * time.Second) // now=8，已过期(8>7)
	if got := c.Get(1); got != -1 {
		t.Fatalf("Get(1) = %d, want -1 (停止访问后最终过期)", got)
	}
}

// 纯惰性策略的关键特性：过期但未被访问的节点不会被主动清理，
// Put 超容量时按 LRU 删尾部，而非优先删过期节点（这是与最小堆版的本质区别）。
func TestTTLLRUCache_NoActiveEvictionOfExpired(t *testing.T) {
	clk := newFakeClock()
	c := NewTTLLRUCache(2, 3*time.Second)
	c.now = clk.now

	c.Put(1, 1)              // expire @3
	clk.add(1 * time.Second) // now=1
	c.Put(2, 2)              // expire @4
	clk.add(5 * time.Second) // now=6：1、2 均已过期，但都未被访问，仍留在缓存中
	c.Put(3, 3)              // 超容量 → 按 LRU 删尾部 1（不会因 2 过期而优先删 2）

	if got := c.Get(1); got != -1 {
		t.Fatalf("Get(1) = %d, want -1 (尾部被 LRU 淘汰)", got)
	}
	if got := c.Get(3); got != 3 {
		t.Fatalf("Get(3) = %d, want 3", got)
	}
	// 2 虽已过期，但纯惰性策略未在 Put(3) 时清理它；直到 Get(2) 命中才被惰性删除
	if got := c.Get(2); got != -1 {
		t.Fatalf("Get(2) = %d, want -1 (过期，命中时惰性删除)", got)
	}
}

// Put 命中已过期 key 会「复活」：重置过期时刻
func TestTTLLRUCache_PutRevivesExpiredKey(t *testing.T) {
	clk := newFakeClock()
	c := NewTTLLRUCache(1, 3*time.Second)
	c.now = clk.now

	c.Put(1, 1)              // expire @3
	clk.add(4 * time.Second) // now=4，1 已过期但仍在 map 中
	c.Put(1, 11)             // 命中（即便已过期）→ 更新为 11 并续期到 4+3=7
	clk.add(2 * time.Second) // now=6，续期后未过期(7>6)
	if got := c.Get(1); got != 11 {
		t.Fatalf("Get(1) = %d, want 11 (过期 key 被 Put 复活)", got)
	}
}
