find . \( -type d -o -type f \) -name "*.sh" | cut -d "." -f2 | 
sed -e 's/\///g'
