#!/bin/bash

# Verifica parámetro
if [ -z "$1" ]; then
    echo "Uso: $0 <n>"
    exit 1
fi

n="$1"
dir="solutions/day$n"
pkg="day$n"

# 1. Crear carpeta y archivos
mkdir -p "$dir"

cat > "$dir/Solution.go" <<EOF
package $pkg

func Solution(input string) (string, string) {
    // TODO: implement
    return "", ""
}
EOF

cat > "$dir/Utils.go" <<EOF
package $pkg

// Utils para el día $n
EOF

cat > "inputs/day$n.txt" <<EOF
EOF

# 2. Insertar import en main.go
# Busca la línea que contiene "import (" y agrega justo después
sed -i "" "/import (/a\\
    \"advent_of_code/solutions/day$n\"\

# 3. Insertar la entrada al map solutions
# Busca el inicio del map: "var solutions = map[string]func(string) (string, string){" 
# Y agrega la línea siguiente
sed -i "" "/var solutions = map\[string\]func(string) (string, string){/a\\
    \"$n\": day$n.Solution,\
