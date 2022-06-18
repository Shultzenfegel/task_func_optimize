package gavrilov_solution

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/comdiv/task_func_optimize_base_go/basis"
)

//BenchmarkBasicSuperFuncImpl оставлен для наглядного сравнения производительности
func BenchmarkBasicSuperFuncImpl(b *testing.B) {
	basis.SuperFuncBenchmark(basis.BasicSuperFuncImpl, b)
}

func TestMySuperFuncImpl(t *testing.T) {
	basis.SuperFuncTestCase(MySuperFuncImpl, t)
}

func BenchmarkGavrilovSuperFuncImpl(b *testing.B) {
	basis.SuperFuncBenchmark(MySuperFuncImpl, b)
}

func TestMySuperFuncImpl2(t *testing.T) {
	basis.SuperFuncTestCase(MySuperFuncImpl2, t)
}

func BenchmarkGavrilovSuperFuncImpl2(b *testing.B) {
	basis.SuperFuncBenchmark(MySuperFuncImpl2, b)
}

func TestMySuperFuncImpl3(t *testing.T) {
	basis.SuperFuncTestCase(MySuperFuncImpl3, t)
}

func BenchmarkGavrilovSuperFuncImpl3(b *testing.B) {
	basis.SuperFuncBenchmark(MySuperFuncImpl3, b)
}

func TestMySuperFuncImpl4(t *testing.T) {
	basis.SuperFuncTestCase(MySuperFuncImpl4, t)
}

func BenchmarkGavrilovSuperFuncImpl4(b *testing.B) {
	basis.SuperFuncBenchmark(MySuperFuncImpl4, b)
}

func TestMySuperFuncImpl5(t *testing.T) {
	basis.SuperFuncTestCase(MySuperFuncImpl5, t)
}

func BenchmarkGavrilovSuperFuncImpl5(b *testing.B) {
	basis.SuperFuncBenchmark(MySuperFuncImpl5, b)
}

func TestFib(t *testing.T) {
	t.Run("`n==0` -> 0", func(t *testing.T) {
		assert.Equal(t, uint32(0), Fib(0))
	})
	t.Run("`n==1` -> 1", func(t *testing.T) {
		assert.Equal(t, uint32(1), Fib(1))
	})
	t.Run("`n==2` -> 1", func(t *testing.T) {
		assert.Equal(t, uint32(1), Fib(2))
	})
	t.Run("`n==3` -> 2", func(t *testing.T) {
		assert.Equal(t, uint32(2), Fib(3))
	})
	t.Run("`n==4` -> 3", func(t *testing.T) {
		assert.Equal(t, uint32(3), Fib(4))
	})
	t.Run("`n==5` -> 5", func(t *testing.T) {
		assert.Equal(t, uint32(5), Fib(5))
	})
	t.Run("`n==6` -> 8", func(t *testing.T) {
		assert.Equal(t, uint32(8), Fib(6))
	})
	t.Run("`n==7` -> 13", func(t *testing.T) {
		assert.Equal(t, uint32(13), Fib(7))
	})
}

func TestPower(t *testing.T) {
	t.Run("`p==0` -> 1", func(t *testing.T) {
		assert.Equal(t, float64(1), Power(2, 0))
	})
	t.Run("`p==1` -> x", func(t *testing.T) {
		assert.Equal(t, float64(3), Power(3, 1))
	})
	t.Run("`p==2` -> x", func(t *testing.T) {
		assert.Equal(t, float64(9), Power(3, 2))
	})
	t.Run("`p==3` -> x", func(t *testing.T) {
		assert.Equal(t, float64(27), Power(3, 3))
	})
	t.Run("`p==4` -> x", func(t *testing.T) {
		assert.Equal(t, float64(81), Power(3, 4))
	})
	t.Run("`p==5` -> x", func(t *testing.T) {
		assert.Equal(t, float64(243), Power(3, 5))
	})
	t.Run("`p==6` -> x", func(t *testing.T) {
		assert.Equal(t, float64(729), Power(3, 6))
	})
	t.Run("`p==7` -> x", func(t *testing.T) {
		assert.Equal(t, float64(2187), Power(3, 7))
	})
	t.Run("`p==8` -> x", func(t *testing.T) {
		assert.Equal(t, float64(6561), Power(3, 8))
	})
	t.Run("`p==9` -> x", func(t *testing.T) {
		assert.Equal(t, float64(19683), Power(3, 9))
	})
	t.Run("`p==10` -> x", func(t *testing.T) {
		assert.Equal(t, float64(59049), Power(3, 10))
	})
}
