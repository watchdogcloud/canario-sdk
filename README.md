# canario
go-sdk for metric collection



- user installs canario 
- user makes a canario.yaml
- puts in what metrics they want to opt in for 
- some metrics would just fetch all the data (like cpu , disk etc )
- after that , metrics would be batched and pushed in kafka mgmt server 
- user can make gauges and counters that would increment and make api requests instantly 
- user can use our logger to push logs to api req