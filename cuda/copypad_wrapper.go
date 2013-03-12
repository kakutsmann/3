package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var copypad_code cu.Function

type copypad_args struct {
	arg_dst unsafe.Pointer
	arg_D0  int
	arg_D1  int
	arg_D2  int
	arg_src unsafe.Pointer
	arg_S0  int
	arg_S1  int
	arg_S2  int
	argptr  [8]unsafe.Pointer
}

// Wrapper for copypad CUDA kernel, asynchronous.
func k_copypad_async(dst unsafe.Pointer, D0 int, D1 int, D2 int, src unsafe.Pointer, S0 int, S1 int, S2 int, cfg *config, str cu.Stream) {
	if copypad_code == 0 {
		copypad_code = fatbinLoad(copypad_map, "copypad")
	}

	var a copypad_args

	a.arg_dst = dst
	a.argptr[0] = unsafe.Pointer(&a.arg_dst)
	a.arg_D0 = D0
	a.argptr[1] = unsafe.Pointer(&a.arg_D0)
	a.arg_D1 = D1
	a.argptr[2] = unsafe.Pointer(&a.arg_D1)
	a.arg_D2 = D2
	a.argptr[3] = unsafe.Pointer(&a.arg_D2)
	a.arg_src = src
	a.argptr[4] = unsafe.Pointer(&a.arg_src)
	a.arg_S0 = S0
	a.argptr[5] = unsafe.Pointer(&a.arg_S0)
	a.arg_S1 = S1
	a.argptr[6] = unsafe.Pointer(&a.arg_S1)
	a.arg_S2 = S2
	a.argptr[7] = unsafe.Pointer(&a.arg_S2)

	args := a.argptr[:]
	cu.LaunchKernel(copypad_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, str, args)
}

// Wrapper for copypad CUDA kernel, synchronized.
func k_copypad(dst unsafe.Pointer, D0 int, D1 int, D2 int, src unsafe.Pointer, S0 int, S1 int, S2 int, cfg *config) {
	str := stream()
	k_copypad_async(dst, D0, D1, D2, src, S0, S1, S2, cfg, str)
	syncAndRecycle(str)
}

var copypad_map = map[int]string{0: ""}
