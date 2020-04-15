declare -a sql_files
sql_files=`ls ./migrations/`

echo "sql_files: ${sql_files}"

for i in ${sql_files[*]}
do
    psql -d fiestadb -f "./migrations/"$i
done 
