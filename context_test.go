package stack

import "testing"

func TestGet(t *testing.T) {
	ctx := NewContext()
	ctx.m["flip"] = "flop"
	ctx.m["bish"] = nil

	val, err := ctx.Get("flip")
	assertEquals(t, nil, err)
	assertEquals(t, "flop", val)

	val, err = ctx.Get("bish")
	assertEquals(t, nil, err)
	assertEquals(t, nil, val)

	_, err = ctx.Get("wibble")
	assertEquals(t, "stack.Context: key \"wibble\" does not exist", err.Error())
}

func TestPut(t *testing.T) {
	ctx := NewContext()

	ctx.Put("bish", "bash")
	assertEquals(t, "bash", ctx.m["bish"])
}

func TestDelete(t *testing.T) {
	ctx := NewContext()
	ctx.m["flip"] = "flop"

	ctx.Delete("flip")
	assertEquals(t, nil, ctx.m["flip"])
}
