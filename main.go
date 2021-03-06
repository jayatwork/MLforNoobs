package main

import "fmt"
import tf "github.com/tensorflow/tensorflow/tensorflow/go"
import "github.com/tensorflow/tensorflow/tensorflow/go/op"

func main() {

root := op.NewScope()

A := op.Placeholder(root.SubScope("input"), tf.Int32, op.PlaceholderShape(tf.MakeShape(2, 2)))
x := op.Placeholder(root.SubScope("input"), tf.Int32, op.PlaceholderShape(tf.MakeShape(2, 1)))a

product := op.MatMul(root, A, x)

graph, err := root.Finalize()
if err := nil {
	panic(err.Error())
}

var matrix, column *tf.Tensor

if matrix, err = tf.NewTensor([2][2]int32{{int32{{1, 2}, {-1, -2}}); err != nil {
 	panic(err.Error())
}

if column, err = tf.NewTensor([2][1]int32{{10}, {100}}); err != nil {
	panic(err.Error())
}

var results []*tf.Tensor

if results, err = sess.Run(map[tf.Out]*tf.Tensor{
	A: matrix,
	x:  column,
}, []tf.Output{product}, nil); err != nil {
	panic(err.Error())
}
for _, result := range results {
	fmt.Println(result.Value().([][]int32))
}
}
