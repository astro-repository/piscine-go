find . \( -type d -o -type f \) -name "*.sh" | sed -e 's/\(\.sh\)$//g' | sed -e 's/^\(\.\/\)//g'
