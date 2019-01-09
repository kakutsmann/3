#! /bin/bash

go build cuda2go.go || exit 1

if [ -z ${CUDA_BIN_PATH+x} ] ; then
    NVCC='nvcc -std c++03 --compiler-options -Werror --compiler-options -Wall -Xptxas -O3 -ptx'
else
    NVCC="${CUDA_BIN_PATH}/nvcc -std c++03 --compiler-options -Werror --compiler-options -Wall -Xptxas -O3 -ptx"
fi
if [ -z ${CUDA_INC_PATH+x} ] ; then CUDA_INC_PATH="/usr/local/cuda/include"; fi

for f in *.cu; do
	g=$(echo $f | sed 's/\.cu$//') # file basename
	for cc in 30 35 37 50 52 53 60 61 70 75; do
		if [[ $f -nt $g'_'$cc.ptx ]]; then
			echo $NVCC -gencode arch=compute_$cc,code=sm_$cc $f -o $g'_'$cc.ptx
			$NVCC -I"$CUDA_INC_PATH" -gencode arch=compute_$cc,code=sm_$cc $f -o $g'_'$cc.ptx # error can be ignored
		fi
	done
	if [[ $f -nt $g'_wrapper.go' ]]; then
		./cuda2go $f || exit 1
	fi
done

