#include <stdint.h>
#include "exchange.h"
#include "float3.h"
#include "stencil.h"

extern "C" __global__ void
adddmibulk(float* __restrict__ Hx, float* __restrict__ Hy, float* __restrict__ Hz,
           float* __restrict__ mx, float* __restrict__ my, float* __restrict__ mz,
           float* __restrict__ aLUT2d,
           float* __restrict__ DxLUT2d, float* __restrict__ DyLUT2d, float* __restrict__ DzLUT2d,
           uint8_t* __restrict__ regions,
           float cx, float cy, float cz, int Nx, int Ny, int Nz, uint8_t PBC) {

	int ix = blockIdx.x * blockDim.x + threadIdx.x;
	int iy = blockIdx.y * blockDim.y + threadIdx.y;
	int iz = blockIdx.z * blockDim.z + threadIdx.z;

	if (ix >= Nx || iy >= Ny || iz >= Nz) {
		return;
	}

	int I = idx(ix, iy, iz);                      // central cell index
	float3 h = make_float3(Hx[I], Hy[I], Hz[I]);  // add to H
	float3 m0 = make_float3(mx[I], my[I], mz[I]); // central m
	uint8_t r0 = regions[I];
	int i_;                                       // neighbor index

	if(is0(m0)) {
		return;
	}

	// x derivatives (along length)
	{
		float3 m1 = make_float3(0.0f, 0.0f, 0.0f);     // left neighbor
		i_ = idx(lclampx(ix-1), iy, iz);               // load neighbor m if inside grid, keep 0 otherwise
		if (ix-1 >= 0 || PBCx) {
			m1 = make_float3(mx[i_], my[i_], mz[i_]);
		}
		float A1 = aLUT2d[symidx(r0, regions[i_])];    // inter-region Aex
		float Dy = DyLUT2d[symidx(r0, regions[i_])];
		float Dz = DzLUT2d[symidx(r0, regions[i_])];
		if (is0(m1)) {                                 // neighbor missing
			A1 = 0;
			Dy = 0;
			Dz = 0;
		}
		h   += (2.0f*A1/(cx*cx)) * (m1 - m0);          // exchange
		h.y -= (Dy/cx)*(m0.z - m1.z);                  // DM (first 1/2 contrib. 2*D * deltaM / (2*c)) !! ?? if boundary, other x2 ?? !!
		h.z += (Dz/cx)*(m0.y - m1.y);
	}


	{
		float3 m2 = make_float3(0.0f, 0.0f, 0.0f);     // right neighbor
		i_ = idx(hclampx(ix+1), iy, iz);
		if (ix+1 < Nx || PBCx) {
			m2 = make_float3(mx[i_], my[i_], mz[i_]);
		}
		float A2 = aLUT2d[symidx(r0, regions[i_])];
		float Dy = DyLUT2d[symidx(r0, regions[i_])];
		float Dz = DzLUT2d[symidx(r0, regions[i_])];
		if (is0(m2)) {
			A2 = 0;
			Dy = 0;
			Dz = 0;
		}
		h   += (2.0f*A2/(cx*cx)) * (m2 - m0);
		h.y -= (Dy/cx)*(m2.z - m0.z);
		h.z += (Dz/cx)*(m2.y - m0.y);
	}

	// y derivatives (along height)
	{
		float3 m1 = make_float3(0.0f, 0.0f, 0.0f);
		i_ = idx(ix, lclampy(iy-1), iz);
		if (iy-1 >= 0 || PBCy) {
			m1 = make_float3(mx[i_], my[i_], mz[i_]);
		}
		float A1 = aLUT2d[symidx(r0, regions[i_])];
		float Dx = DxLUT2d[symidx(r0, regions[i_])];
		float Dz = DzLUT2d[symidx(r0, regions[i_])];
		if (is0(m1)) {
			A1 = 0;
			Dx = 0;
			Dz = 0;
		}
		h   += (2.0f*A1/(cy*cy)) * (m1 - m0);
		h.x += (Dx/cy)*(m0.z - m1.z);
		h.z -= (Dz/cy)*(m0.x - m1.x);
	}

	{
		float3 m2 = make_float3(0.0f, 0.0f, 0.0f);
		i_ = idx(ix, hclampy(iy+1), iz);
		if  (iy+1 < Ny || PBCy) {
			m2 = make_float3(mx[i_], my[i_], mz[i_]);
		}
		float A2 = aLUT2d[symidx(r0, regions[i_])];
		float Dx = DxLUT2d[symidx(r0, regions[i_])];
		float Dz = DzLUT2d[symidx(r0, regions[i_])];
		if (is0(m2)) {
			A2 = 0;
			Dx = 0;
			Dz = 0;
		}
		h   += (2.0f*A2/(cy*cy)) * (m2 - m0);
		h.x += (Dx/cy)*(m2.z - m0.z);
		h.z -= (Dz/cy)*(m2.x - m0.x);
	}

	// only take vertical derivative for 3D sim
	if (Nz != 1) {
		//	// bottom neighbor
		//	{
		//		i_  = idx(ix, iy, lclampz(iz-1));
		//		float3 m1  = make_float3(mx[i_], my[i_], mz[i_]);
		//		m1  = ( is0(m1)? m0: m1 );                         // Neumann BC
		//		float A1 = aLUT2d[symidx(r0, regions[i_])];
		//		h += (2.0f*A1/(cz*cz)) * (m1 - m0);                // Exchange only
		//	}

		//	// top neighbor
		//	{
		//		i_  = idx(ix, iy, hclampz(iz+1));
		//		float3 m2  = make_float3(mx[i_], my[i_], mz[i_]);
		//		m2  = ( is0(m2)? m0: m2 );
		//		float A2 = aLUT2d[symidx(r0, regions[i_])];
		//		h += (2.0f*A2/(cz*cz)) * (m2 - m0);
		//	}
	}

	// write back, result is H + Hdmi + Hex
	Hx[I] = h.x;
	Hy[I] = h.y;
	Hz[I] = h.z;
}

// Note on boundary conditions.
//
// We need the derivative and laplacian of m in point A, but e.g. C lies out of the boundaries.
// We use the boundary condition in B (derivative of the magnetization) to extrapolate m to point C:
// 	m_C = m_A + (dm/dx)|_B * cellsize
//
// When point C is inside the boundary, we just use its actual value.
//
// Then we can take the central derivative in A:
// 	(dm/dx)|_A = (m_C - m_D) / (2*cellsize)
// And the laplacian:
// 	lapl(m)|_A = (m_C + m_D - 2*m_A) / (cellsize^2)
//
// All these operations should be second order as they involve only central derivatives.
//
//    ------------------------------------------------------------------ *
//   |                                                   |             C |
//   |                                                   |          **   |
//   |                                                   |        ***    |
//   |                                                   |     ***       |
//   |                                                   |   ***         |
//   |                                                   | ***           |
//   |                                                   B               |
//   |                                               *** |               |
//   |                                            ***    |               |
//   |                                         ****      |               |
//   |                                     ****          |               |
//   |                                  ****             |               |
//   |                              ** A                 |               |
//   |                         *****                     |               |
//   |                   ******                          |               |
//   |          *********                                |               |
//   |D ********                                         |               |
//   |                                                   |               |
//   +----------------+----------------+-----------------+---------------+
//  -1              -0.5               0               0.5               1
//                                 x
