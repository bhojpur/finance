# Bhojpur Finance - Linear & Time-Invariant

```go get github.com/bhojpur/finance```

* State-space representation and estimation of linear, Time-Invariant systems for control theory

	```math
	 x'(t) = A * x(t) + B * u(t)
	```
	 and
	```math
	 y(t)  = C * x(t) + D * u(t)
	```

* Can be used as an input for a Kalman filter. 

## Simple Usage

```go
	// define time-continuous linear system
	system, err := lti.NewSystem(
		...
	)

	// check system properties
	fmt.Println("Observable=", system.MustObservable())
	fmt.Println("Controllable=", system.MustControllable())

	// define initial state (x) and control (u) vectors
	...

	// get derivative vector for new state
	fmt.Println(system.Derivative(x, u))

	// get output vector for new state
	fmt.Println(system.Response(x, u))

	// discretize LTI system and propagate state by time step dt
	discrete, err := system.Discretize(dt)

	fmt.Println("x(k+1)=", discrete.Predict(x, u))
}
```
