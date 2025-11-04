#!/bin/bash

echo "Ejecutando todas las pruebas de AlgoGram..."
echo "============================================"

cd pruebas_algogram
total=0
pasados=0

for test in $(ls *_in | sed 's/_in//' | sort); do
    total=$((total + 1))
    echo -n "Probando caso $test... "

    if ../algogram ${test}_usuarios < ${test}_in > /tmp/test_${test}_output.txt 2>&1 && \
       diff -q /tmp/test_${test}_output.txt ${test}_out > /dev/null 2>&1; then
        echo "✓ OK"
        pasados=$((pasados + 1))
    else
        echo "✗ FAIL"
    fi
done

echo "============================================"
echo "Resultado: $pasados/$total pruebas pasadas"

if [ $pasados -eq $total ]; then
    echo "¡Todas las pruebas pasaron exitosamente!"
    exit 0
else
    echo "Algunas pruebas fallaron."
    exit 1
fi
