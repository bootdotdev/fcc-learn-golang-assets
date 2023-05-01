# Interfaces are implemented implicitly

A type implements an interface by implementing its methods. Unlike in many other languages, there is no explicit declaration of intent, there is no "implements" keyword.

Implicit interfaces *decouple* the definition of an interface from its implementation. You may add methods to a type and in the process be unknowingly implementing various interfaces, and *that's okay*.
