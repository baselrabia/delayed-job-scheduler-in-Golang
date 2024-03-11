# Benchmark Delayed Job Scheduler
Benchmark an implementation of a scheduler that can efficiently take jobs with a delay time associated with them and execute these jobs after the said amount of delay in seconds.


## Usage

Once installed, you can run the benchmark using the following command:

```bash
go test -bench=. ./...
```

This command will execute the benchmark tests defined in the application and display the results, including performance metrics such as time per operation, memory usage, and allocations.

## Benchmark Result

After running the benchmark, you'll see output similar to the following:

```plaintext
BenchmarkScheduler-8   	Tick at  2024-03-12 00:21:31.183710357 +0200 EET m=+1.001166484
Tick at  2024-03-12 00:21:31.183977618 +0200 EET m=+1.001433749
Tick at  2024-03-12 00:21:31.184834977 +0200 EET m=+1.002291093
Tick at  2024-03-12 00:21:31.19004103 +0200 EET m=+1.007497145
Tick at  2024-03-12 00:21:31.637591175 +0200 EET m=+1.455047303
 3069178	       434.9 ns/op	     224 B/op	       3 allocs/op
PASS
ok  	github.com/baselrabia/delayed-job-scheduler/pkg/delayedjob	1.805s
```

This output provides insights into the performance of the BenchmarkScheduler application, including the number of iterations, time per operation, memory usage, and allocations.

 