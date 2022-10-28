# Mô tả
At the same time, you could be using the same technique to create many types of objects. For example, you'll use almost the same technique to build a car as you would build a bus, except that they'll be of different sizes and number of seats, so why don't we reuse the construction process? This is where the Builder pattern comes to the rescue.

# Mục tiêu
Abstract complex creations so that object creation is separated from the object user Create an object step by step by filling its fields and creating the embedded
objects Reuse the object creation algorithm between many objects

# Example
I must have a manufacturing type that constructs everything that a vehicle needs

When using a car builder, the VehicleProduct with four wheels, five seats, and a structure defined as Car must be returned

When using a motorbike builder, the VehicleProduct with two wheels, two seats, and a structure defined as Motorbike must be returned

A VehicleProduct built by any BuildProcess builder must be open to modifications
