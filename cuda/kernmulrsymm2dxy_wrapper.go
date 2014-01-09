package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"sync"
	"unsafe"
)

// CUDA handle for kernmulRSymm2Dxy kernel
var kernmulRSymm2Dxy_code cu.Function

// Stores the arguments for kernmulRSymm2Dxy kernel invocation
type kernmulRSymm2Dxy_args_t struct {
	arg_fftMx  unsafe.Pointer
	arg_fftMy  unsafe.Pointer
	arg_fftKxx unsafe.Pointer
	arg_fftKyy unsafe.Pointer
	arg_fftKxy unsafe.Pointer
	arg_Nx     int
	arg_Ny     int
	argptr     [7]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for kernmulRSymm2Dxy kernel invocation
var kernmulRSymm2Dxy_args kernmulRSymm2Dxy_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	kernmulRSymm2Dxy_args.argptr[0] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftMx)
	kernmulRSymm2Dxy_args.argptr[1] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftMy)
	kernmulRSymm2Dxy_args.argptr[2] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKxx)
	kernmulRSymm2Dxy_args.argptr[3] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKyy)
	kernmulRSymm2Dxy_args.argptr[4] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_fftKxy)
	kernmulRSymm2Dxy_args.argptr[5] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_Nx)
	kernmulRSymm2Dxy_args.argptr[6] = unsafe.Pointer(&kernmulRSymm2Dxy_args.arg_Ny)
}

// Wrapper for kernmulRSymm2Dxy CUDA kernel, asynchronous.
func k_kernmulRSymm2Dxy_async(fftMx unsafe.Pointer, fftMy unsafe.Pointer, fftKxx unsafe.Pointer, fftKyy unsafe.Pointer, fftKxy unsafe.Pointer, Nx int, Ny int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	kernmulRSymm2Dxy_args.Lock()
	defer kernmulRSymm2Dxy_args.Unlock()

	if kernmulRSymm2Dxy_code == 0 {
		kernmulRSymm2Dxy_code = fatbinLoad(kernmulRSymm2Dxy_map, "kernmulRSymm2Dxy")
	}

	kernmulRSymm2Dxy_args.arg_fftMx = fftMx
	kernmulRSymm2Dxy_args.arg_fftMy = fftMy
	kernmulRSymm2Dxy_args.arg_fftKxx = fftKxx
	kernmulRSymm2Dxy_args.arg_fftKyy = fftKyy
	kernmulRSymm2Dxy_args.arg_fftKxy = fftKxy
	kernmulRSymm2Dxy_args.arg_Nx = Nx
	kernmulRSymm2Dxy_args.arg_Ny = Ny

	args := kernmulRSymm2Dxy_args.argptr[:]
	cu.LaunchKernel(kernmulRSymm2Dxy_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for kernmulRSymm2Dxy kernel.
var kernmulRSymm2Dxy_map = map[int]string{0: "",
	20: kernmulRSymm2Dxy_ptx_20,
	30: kernmulRSymm2Dxy_ptx_30,
	35: kernmulRSymm2Dxy_ptx_35}

// kernmulRSymm2Dxy PTX code for various compute capabilities.
const (
	kernmulRSymm2Dxy_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<13>;
	.reg .f32 	%f<16>;
	.reg .s64 	%rd<18>;


	ld.param.u64 	%rd6, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd7, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd8, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd9, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd10, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd8;
	cvta.to.global.u64 	%rd4, %rd7;
	cvta.to.global.u64 	%rd5, %rd6;
	.loc 1 29 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	.loc 1 30 1
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	.loc 1 32 1
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p2, %p1;
	.loc 1 32 1
	@%p3 bra 	BB0_2;

	.loc 1 37 1
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	.loc 1 38 1
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd11, %r12, 4;
	add.s64 	%rd12, %rd5, %rd11;
	.loc 1 41 1
	ld.global.f32 	%f1, [%rd12+4];
	add.s64 	%rd13, %rd4, %rd11;
	.loc 1 43 1
	ld.global.f32 	%f2, [%rd13+4];
	mul.wide.s32 	%rd14, %r11, 4;
	add.s64 	%rd15, %rd3, %rd14;
	add.s64 	%rd16, %rd2, %rd14;
	.loc 1 48 1
	ld.global.f32 	%f3, [%rd16];
	add.s64 	%rd17, %rd1, %rd14;
	.loc 1 47 1
	ld.global.f32 	%f4, [%rd15];
	.loc 1 40 1
	ld.global.f32 	%f5, [%rd12];
	.loc 1 49 1
	ld.global.f32 	%f6, [%rd17];
	.loc 1 42 1
	ld.global.f32 	%f7, [%rd13];
	.loc 1 57 1
	mul.f32 	%f8, %f7, %f6;
	fma.rn.f32 	%f9, %f5, %f4, %f8;
	st.global.f32 	[%rd12], %f9;
	.loc 1 58 1
	mul.f32 	%f10, %f2, %f6;
	fma.rn.f32 	%f11, %f1, %f4, %f10;
	st.global.f32 	[%rd12+4], %f11;
	.loc 1 59 1
	mul.f32 	%f12, %f7, %f3;
	fma.rn.f32 	%f13, %f5, %f6, %f12;
	st.global.f32 	[%rd13], %f13;
	.loc 1 60 1
	mul.f32 	%f14, %f2, %f3;
	fma.rn.f32 	%f15, %f1, %f6, %f14;
	st.global.f32 	[%rd13+4], %f15;

BB0_2:
	.loc 1 61 2
	ret;
}


`
	kernmulRSymm2Dxy_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<13>;
	.reg .f32 	%f<16>;
	.reg .s64 	%rd<18>;


	ld.param.u64 	%rd6, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd7, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd8, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd9, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd10, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd8;
	cvta.to.global.u64 	%rd4, %rd7;
	cvta.to.global.u64 	%rd5, %rd6;
	.loc 1 29 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	.loc 1 30 1
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	.loc 1 32 1
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p2, %p1;
	.loc 1 32 1
	@%p3 bra 	BB0_2;

	.loc 1 37 1
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	.loc 1 38 1
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd11, %r12, 4;
	add.s64 	%rd12, %rd5, %rd11;
	.loc 1 41 1
	ld.global.f32 	%f1, [%rd12+4];
	add.s64 	%rd13, %rd4, %rd11;
	.loc 1 43 1
	ld.global.f32 	%f2, [%rd13+4];
	mul.wide.s32 	%rd14, %r11, 4;
	add.s64 	%rd15, %rd3, %rd14;
	add.s64 	%rd16, %rd2, %rd14;
	.loc 1 48 1
	ld.global.f32 	%f3, [%rd16];
	add.s64 	%rd17, %rd1, %rd14;
	.loc 1 47 1
	ld.global.f32 	%f4, [%rd15];
	.loc 1 40 1
	ld.global.f32 	%f5, [%rd12];
	.loc 1 49 1
	ld.global.f32 	%f6, [%rd17];
	.loc 1 42 1
	ld.global.f32 	%f7, [%rd13];
	.loc 1 57 1
	mul.f32 	%f8, %f7, %f6;
	fma.rn.f32 	%f9, %f5, %f4, %f8;
	st.global.f32 	[%rd12], %f9;
	.loc 1 58 1
	mul.f32 	%f10, %f2, %f6;
	fma.rn.f32 	%f11, %f1, %f4, %f10;
	st.global.f32 	[%rd12+4], %f11;
	.loc 1 59 1
	mul.f32 	%f12, %f7, %f3;
	fma.rn.f32 	%f13, %f5, %f6, %f12;
	st.global.f32 	[%rd13], %f13;
	.loc 1 60 1
	mul.f32 	%f14, %f2, %f3;
	fma.rn.f32 	%f15, %f1, %f6, %f14;
	st.global.f32 	[%rd13+4], %f15;

BB0_2:
	.loc 1 61 2
	ret;
}


`
	kernmulRSymm2Dxy_ptx_35 = `
.version 3.2
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry kernmulRSymm2Dxy(
	.param .u64 kernmulRSymm2Dxy_param_0,
	.param .u64 kernmulRSymm2Dxy_param_1,
	.param .u64 kernmulRSymm2Dxy_param_2,
	.param .u64 kernmulRSymm2Dxy_param_3,
	.param .u64 kernmulRSymm2Dxy_param_4,
	.param .u32 kernmulRSymm2Dxy_param_5,
	.param .u32 kernmulRSymm2Dxy_param_6
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<13>;
	.reg .f32 	%f<16>;
	.reg .s64 	%rd<18>;


	ld.param.u64 	%rd6, [kernmulRSymm2Dxy_param_0];
	ld.param.u64 	%rd7, [kernmulRSymm2Dxy_param_1];
	ld.param.u64 	%rd8, [kernmulRSymm2Dxy_param_2];
	ld.param.u64 	%rd9, [kernmulRSymm2Dxy_param_3];
	ld.param.u64 	%rd10, [kernmulRSymm2Dxy_param_4];
	ld.param.u32 	%r3, [kernmulRSymm2Dxy_param_5];
	ld.param.u32 	%r4, [kernmulRSymm2Dxy_param_6];
	cvta.to.global.u64 	%rd1, %rd10;
	cvta.to.global.u64 	%rd2, %rd9;
	cvta.to.global.u64 	%rd3, %rd8;
	cvta.to.global.u64 	%rd4, %rd7;
	cvta.to.global.u64 	%rd5, %rd6;
	.loc 1 29 1
	mov.u32 	%r5, %ntid.x;
	mov.u32 	%r6, %ctaid.x;
	mov.u32 	%r7, %tid.x;
	mad.lo.s32 	%r1, %r5, %r6, %r7;
	.loc 1 30 1
	mov.u32 	%r8, %ntid.y;
	mov.u32 	%r9, %ctaid.y;
	mov.u32 	%r10, %tid.y;
	mad.lo.s32 	%r2, %r8, %r9, %r10;
	.loc 1 32 1
	setp.ge.s32	%p1, %r2, %r4;
	setp.ge.s32	%p2, %r1, %r3;
	or.pred  	%p3, %p2, %p1;
	.loc 1 32 1
	@%p3 bra 	BB2_2;

	.loc 1 37 1
	mad.lo.s32 	%r11, %r2, %r3, %r1;
	.loc 1 38 1
	shl.b32 	%r12, %r11, 1;
	mul.wide.s32 	%rd11, %r12, 4;
	add.s64 	%rd12, %rd5, %rd11;
	.loc 1 41 1
	ld.global.f32 	%f1, [%rd12+4];
	add.s64 	%rd13, %rd4, %rd11;
	.loc 1 43 1
	ld.global.f32 	%f2, [%rd13+4];
	mul.wide.s32 	%rd14, %r11, 4;
	add.s64 	%rd15, %rd3, %rd14;
	add.s64 	%rd16, %rd2, %rd14;
	.loc 1 48 1
	ld.global.nc.f32 	%f3, [%rd16];
	add.s64 	%rd17, %rd1, %rd14;
	.loc 1 47 1
	ld.global.nc.f32 	%f4, [%rd15];
	.loc 1 40 1
	ld.global.f32 	%f5, [%rd12];
	.loc 1 49 1
	ld.global.nc.f32 	%f6, [%rd17];
	.loc 1 42 1
	ld.global.f32 	%f7, [%rd13];
	.loc 1 57 1
	mul.f32 	%f8, %f7, %f6;
	fma.rn.f32 	%f9, %f5, %f4, %f8;
	st.global.f32 	[%rd12], %f9;
	.loc 1 58 1
	mul.f32 	%f10, %f2, %f6;
	fma.rn.f32 	%f11, %f1, %f4, %f10;
	st.global.f32 	[%rd12+4], %f11;
	.loc 1 59 1
	mul.f32 	%f12, %f7, %f3;
	fma.rn.f32 	%f13, %f5, %f6, %f12;
	st.global.f32 	[%rd13], %f13;
	.loc 1 60 1
	mul.f32 	%f14, %f2, %f3;
	fma.rn.f32 	%f15, %f1, %f6, %f14;
	st.global.f32 	[%rd13+4], %f15;

BB2_2:
	.loc 1 61 2
	ret;
}


`
)