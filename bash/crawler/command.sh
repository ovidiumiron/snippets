 wget -qO- https://gist.githubusercontent.com/ovidiumiron/d25d31b414d859dc1d93df9fba60e8bb/raw/a6dc48f5feb05ae35d887b588b33171262c0ecd0/gistfile1.txt |  cut -d, -f3 | sort | uniq -c | sort -nr | head -n 5
