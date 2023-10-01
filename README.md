# MSDS 431 - Week 2

### Setup (Windows 11 + Git Bash)
- Clone repo with `git clone git@github.com:jeremycruzz/msds301-wk2.git`
- Install go dependencies `go mod tidy`
- Install py dependencies `py -m pip install pandas numpy statsmodels`

### Building executable
- Run `go build src/app/main.go`

### Running Go executable
- Run `main.exe`

### Running Tests
- Run tests with `go test -v ./...`

### Running benchmark
- run `chmod +x benchmark.sh`
- run `./benchmark.sh {RUNS}` where {RUNS} is the amount of runs you want.

<details>
<summary> Results from python (coefficients) </summary>

```
               coef    std err          t      P>|t|      [0.025      0.975]
------------------------------------------------------------------------------
const          3.0001      1.125      2.667      0.026       0.456       5.544
x1             0.5001      0.118      4.241      0.002       0.233       0.767
------------------------------------------------------------------------------
const          3.0009      1.125      2.667      0.026       0.455       5.547
x2             0.5000      0.118      4.239      0.002       0.233       0.767
------------------------------------------------------------------------------
const          3.0025      1.124      2.670      0.026       0.459       5.546
x3             0.4997      0.118      4.239      0.002       0.233       0.766
------------------------------------------------------------------------------
const          3.0017      1.124      2.671      0.026       0.459       5.544
x4             0.4999      0.118      4.243      0.002       0.233       0.766
```

</details>

<details>
<summary>Results from python (full)</summary>
  
```
                            OLS Regression Results                            
==============================================================================
Dep. Variable:                     y1   R-squared:                       0.667
Model:                            OLS   Adj. R-squared:                  0.629
Method:                 Least Squares   F-statistic:                     17.99
Date:                Wed, 20 Sep 2023   Prob (F-statistic):            0.00217
Time:                        01:36:06   Log-Likelihood:                -16.841
No. Observations:                  11   AIC:                             37.68
Df Residuals:                       9   BIC:                             38.48
Df Model:                           1                                         
Covariance Type:            nonrobust                                         
==============================================================================
                 coef    std err          t      P>|t|      [0.025      0.975]
------------------------------------------------------------------------------
const          3.0001      1.125      2.667      0.026       0.456       5.544
x1             0.5001      0.118      4.241      0.002       0.233       0.767
==============================================================================
Omnibus:                        0.082   Durbin-Watson:                   3.212
Prob(Omnibus):                  0.960   Jarque-Bera (JB):                0.289
Skew:                          -0.122   Prob(JB):                        0.865
Kurtosis:                       2.244   Cond. No.                         29.1
==============================================================================

Notes:
[1] Standard Errors assume that the covariance matrix of the errors is correctly specified.
                            OLS Regression Results                            
==============================================================================
Dep. Variable:                     y2   R-squared:                       0.666
Model:                            OLS   Adj. R-squared:                  0.629
Method:                 Least Squares   F-statistic:                     17.97
Date:                Wed, 20 Sep 2023   Prob (F-statistic):            0.00218
Time:                        01:36:06   Log-Likelihood:                -16.846
No. Observations:                  11   AIC:                             37.69
Df Residuals:                       9   BIC:                             38.49
Df Model:                           1                                         
Covariance Type:            nonrobust                                         
==============================================================================
                 coef    std err          t      P>|t|      [0.025      0.975]
------------------------------------------------------------------------------
const          3.0009      1.125      2.667      0.026       0.455       5.547
x2             0.5000      0.118      4.239      0.002       0.233       0.767
==============================================================================
Omnibus:                        1.594   Durbin-Watson:                   2.188
Prob(Omnibus):                  0.451   Jarque-Bera (JB):                1.108
Skew:                          -0.567   Prob(JB):                        0.575
Kurtosis:                       1.936   Cond. No.                         29.1
==============================================================================

Notes:
[1] Standard Errors assume that the covariance matrix of the errors is correctly specified.
                            OLS Regression Results                            
==============================================================================
Dep. Variable:                     y3   R-squared:                       0.666
Model:                            OLS   Adj. R-squared:                  0.629
Method:                 Least Squares   F-statistic:                     17.97
Date:                Wed, 20 Sep 2023   Prob (F-statistic):            0.00218
Time:                        01:36:06   Log-Likelihood:                -16.838
No. Observations:                  11   AIC:                             37.68
Df Residuals:                       9   BIC:                             38.47
Df Model:                           1                                         
Covariance Type:            nonrobust                                         
==============================================================================
                 coef    std err          t      P>|t|      [0.025      0.975]
------------------------------------------------------------------------------
const          3.0025      1.124      2.670      0.026       0.459       5.546
x3             0.4997      0.118      4.239      0.002       0.233       0.766
==============================================================================
Omnibus:                       19.540   Durbin-Watson:                   2.144
Prob(Omnibus):                  0.000   Jarque-Bera (JB):               13.478
Skew:                           2.041   Prob(JB):                      0.00118
Kurtosis:                       6.571   Cond. No.                         29.1
==============================================================================

Notes:
[1] Standard Errors assume that the covariance matrix of the errors is correctly specified.
                            OLS Regression Results                            
==============================================================================
Dep. Variable:                     y4   R-squared:                       0.667
Model:                            OLS   Adj. R-squared:                  0.630
Method:                 Least Squares   F-statistic:                     18.00
Date:                Wed, 20 Sep 2023   Prob (F-statistic):            0.00216
Time:                        01:36:06   Log-Likelihood:                -16.833
No. Observations:                  11   AIC:                             37.67
Df Residuals:                       9   BIC:                             38.46
Df Model:                           1                                         
Covariance Type:            nonrobust                                         
==============================================================================
                 coef    std err          t      P>|t|      [0.025      0.975]
------------------------------------------------------------------------------
const          3.0017      1.124      2.671      0.026       0.459       5.544
x4             0.4999      0.118      4.243      0.002       0.233       0.766
==============================================================================
Omnibus:                        0.555   Durbin-Watson:                   1.662
Prob(Omnibus):                  0.758   Jarque-Bera (JB):                0.524
Skew:                           0.010   Prob(JB):                        0.769
Kurtosis:                       1.931   Cond. No.                         29.1
==============================================================================

Notes:
[1] Standard Errors assume that the covariance matrix of the errors is correctly specified.
```
</details>

### Test Results
The stats package produces the correct linear regression `m` and `b` coefficients using the Anscombe Quartet data set within a delta of `.001`.

### Results

|                 | Warmup 1 (50 runs) | Warmup 2 (100 runs) | Benchmark (1000 runs) |
|-----------------|--------------------:|--------------------:|----------------------:|
| Average time for Go     |       94,144,852 ns |        75,639,160 ns |          83,543,659 ns |
| Average time for Python |      786,503,840 ns |       727,665,684 ns |         732,136,137 ns |
| Average time for R      |      316,122,530 ns |       271,286,936 ns |         288,139,136 ns |

### Analysis
With our benchmark we can see that go is **8.76** times faster than python and **3.45** times faster than R. This highlights the advantage of using Go vs Python or R in our data processing pipeline. While the results follow the results of previous tests, the results of this test is does not produce as extreme values. This could be due to a few reasons:

- The stats package in Go did not give `m` and `b` coeffecients directly, rather it gave points on the linear regression line. A function was written to produce the `m` and `b` coeffecients and may not be the most optimized.

- Benchmarking was done using a `bash` script rather than having each program use their own language's benchmark utilities.

- The test removed the plotting from all languages. Others may have kept them

### Conclusion
Go and the stats package provided sufficient (delta < .001) linear regression coeffecients results using the Anscombe Quartet dataset. The Go program ran much faster than the python and R for the Anscombe Quartet dataset and will likely be even faster for larger datasets.

There are a few downsides to switching to go for our data processing pipeline: 
- Existing programs in python will need to be re-written in go. 
- The package doesn't give us the linear regression coefficients without extra work.
- The learning curve for our developers and analysts to become comfortable with go.

While existing programs will need to eventually be ported into go, we could write all future programs in go and not change existing programs until either a bottle-neck is found or time allows us to. The package doesn't give us the exact data we need but it is relatively easy to produce the data we need with additional programming. Go isn't too hard to learn. The time we lose for developers and analysts to 'ramp up' with go is worth the speed we are getting with our data. 

With all of these points in mind I strongly reccomend that we use Go as our primary language moving forward.