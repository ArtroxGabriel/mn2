# Numerical Differentiation

Numerical differentiation is a technique used to estimate the derivative of a function at a point using function values at discrete points. This is particularly useful when the function is given as a set of data points or when the analytical derivative is difficult or impossible to compute.

## Finite Difference Methods

Finite difference methods approximate derivatives by replacing them with differences between function values at nearby points. The choice of which points to use determines the specific method.

### Forward Difference

The forward difference method approximates the derivative at a point `x` using the function value at `x` and `x+h`, where `h` is a small step size.

**Formula:**

```
f'(x) ≈ (f(x+h) - f(x)) / h
```

**Explanation:**
This method uses the slope of the line connecting the points `(x, f(x))` and `(x+h, f(x+h))` to approximate the tangent at `x`. It is a first-order method.

### Backward Difference

The backward difference method approximates the derivative at a point `x` using the function value at `x` and `x-h`.

**Formula:**

```
f'(x) ≈ (f(x) - f(x-h)) / h
```

**Explanation:**
This method uses the slope of the line connecting the points `(x-h, f(x-h))` and `(x, f(x))` to approximate the tangent at `x`. It is also a first-order method.

### Central Difference

The central difference method approximates the derivative at a point `x` using the function values at `x-h` and `x+h`.

**Formula:**

```
f'(x) ≈ (f(x+h) - f(x-h)) / (2h)
```

**Explanation:**
This method uses the slope of the line connecting the points `(x-h, f(x-h))` and `(x+h, f(x+h))`. It generally provides a more accurate approximation than forward or backward difference methods for the same step size `h` because it considers information symmetrically around `x`. It is a second-order method for approximating the first derivative.

## Order of Accuracy

The order of accuracy of a finite difference method indicates how the error of the approximation changes as the step size `h` is reduced.

- **O(h) (First-order accuracy):** The error is approximately proportional to `h`. Halving `h` roughly halves the error. Forward and backward difference methods are typically O(h).
- **O(h^2) (Second-order accuracy):** The error is approximately proportional to `h^2`. Halving `h` roughly quarters the error. The central difference method for the first derivative is typically O(h^2).
- **O(h^4) (Fourth-order accuracy):** The error is approximately proportional to `h^4`. Halving `h` reduces the error by a factor of about 16. Higher-order methods can be derived by using more points, such as the five-point stencil for the first derivative, which is O(h^4).

Higher-order methods generally provide better accuracy for a given step size `h`, but they may require more function evaluations and can be more susceptible to round-off errors if `h` is too small.

## Convergence Analysis

Convergence analysis in numerical differentiation examines how the approximation approaches the true derivative as the step size `h` tends to zero.

Ideally, as `h` decreases, the truncation error (the error from approximating the derivative with a finite difference formula) also decreases. For example, for an O(h^2) method, the truncation error is proportional to `h^2`. So, reducing `h` should improve accuracy.

However, there's a trade-off. When `h` becomes very small, round-off errors due to the limited precision of computer arithmetic can become significant. Subtracting two very close numbers (like `f(x+h)` and `f(x)`) can lead to a loss of precision, and then dividing by a very small `h` can amplify this error.

Therefore, there is an optimal value of `h` that balances truncation error and round-off error. Choosing an excessively small `h` can lead to less accurate results than a moderately small `h`. Techniques like Richardson extrapolation can be used to estimate the derivative more accurately and also to estimate the optimal `h`.

# Numerical Integration (Quadrature)

Numerical integration, also known as quadrature, involves approximating the value of a definite integral. This is essential when the antiderivative of the integrand is unknown or difficult to find, or when the function is only known at discrete data points.

## Newton-Cotes Formulas

Newton-Cotes formulas are a group of numerical integration rules based on evaluating the integrand at equally spaced points. The general idea is to approximate the function to be integrated by an interpolating polynomial of a certain degree and then integrate this polynomial.

- **General Idea:** Replace the function `f(x)` over the interval `[a, b]` with a polynomial that is easy to integrate. The degree of the polynomial and the points used for interpolation determine the specific rule.
- **Open vs. Closed Formulas:**
  - **Closed formulas** use the function values at the endpoints of the integration interval. Examples include the Trapezoidal rule and Simpson's rules.
  - **Open formulas** use only function values at points strictly within the integration interval. These are useful for integrands with singularities at the endpoints. An example is the Midpoint rule (which is also the simplest open Newton-Cotes formula, based on a zero-degree polynomial or constant).

### Common Newton-Cotes Methods

- **Trapezoidal Rule:** This rule approximates the integral by fitting a first-degree polynomial (a straight line) between the function values at the endpoints of the interval (or each subinterval in the composite version). The area under this line (a trapezoid) approximates the integral.
- **Simpson's Rule (1/3 Rule):** This rule uses a second-degree polynomial (a parabola) to interpolate the function, using three equally spaced points: the two endpoints and the midpoint of the interval. It generally provides higher accuracy than the Trapezoidal rule for smooth functions.
- **Simpson's Rule (3/8 Rule):** This rule uses a third-degree polynomial to interpolate the function, using four equally spaced points. It can be more accurate than the 1/3 rule for some functions but requires one more function evaluation.

## Gaussian Quadrature

Gaussian quadrature formulas offer an alternative approach that often achieves higher accuracy for the same number of function evaluations compared to Newton-Cotes rules.

- **General Idea:** Instead of fixing the abscissas (x-values) to be equally spaced, Gaussian quadrature methods choose the locations of the evaluation points (nodes) and the weights optimally. These nodes are typically the roots of a family of orthogonal polynomials. By choosing these nodes and weights strategically, Gaussian quadrature can exactly integrate polynomials of degree `2n-1` with only `n` function evaluations.
- **Gauss-Legendre Quadrature:** This is a common type of Gaussian quadrature used for integrals over the interval `[-1, 1]`. The nodes are the roots of Legendre polynomials, and the weights are chosen to achieve the highest possible accuracy. Integrals over other intervals `[a, b]` can be transformed to `[-1, 1]` using a linear change of variable.

## Convergence and Error in Numerical Integration

Similar to numerical differentiation, the error in numerical integration depends on the method used and the number of evaluation points (or the width of subintervals, `h`, in composite rules).

- **Truncation Error:** This error arises from approximating the function with a simpler one (e.g., a polynomial). For Newton-Cotes rules like the composite Trapezoidal rule, the error is typically of order O(h^2), while for the composite Simpson's 1/3 rule, it's O(h^4), where `h` is the width of the subintervals. Gaussian quadrature with `n` points can integrate polynomials up to degree `2n-1` exactly, leading to very rapid convergence for smooth functions.
- **Round-off Error:** As with differentiation, performing many calculations with finite precision can lead to an accumulation of round-off errors. However, in integration, round-off errors are generally less problematic than in differentiation because the operations involved (summation and multiplication by weights) are less sensitive to small `h` values than division by `h` or `h^2`.

Generally, increasing the number of evaluation points (or decreasing `h` in composite rules) improves the accuracy of the approximation by reducing the truncation error. However, for very large numbers of points, round-off error might eventually start to increase, though this is less of a concern than in differentiation. Adaptive quadrature methods adjust the step size `h` (or the number of points) in different parts of the integration domain to achieve a desired level of accuracy efficiently.
