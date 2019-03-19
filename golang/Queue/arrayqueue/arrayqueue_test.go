package arrayqueue_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"Algorithm-Datastruct/golang/Queue/arrayqueue"
)
var testarrayqueue *arrayqueue.ArrayQueue

var _ = Describe("Arrayqueue", func() {
	BeforeEach(func() {
		testarrayqueue = arrayqueue.NewArrayQueue(5)

	})
	Describe("Generate a queue",func(){
		It("Should return a queue", func() {
			Expect(testarrayqueue.Candequeue()).To(Equal(false))
			Expect(testarrayqueue.Data).To(Equal(make([]interface{},5,5)))
		})
	})
	Describe("Now Let us enqueue",func(){
		It("enqueue to a queue", func() {
			for i:=1;i<=5;i++{
				testarrayqueue.Enqueue(i)
			}
			testvalue := make([]interface{},5,5)
			for i:=1;i<=5;i++{
				testvalue[i-1] = i
			}
			Expect(testarrayqueue.Candequeue()).To(Equal(true))
			Expect(testarrayqueue.Data).To(Equal(testvalue))
		})
	})
	Describe("Now Let dequeue",func(){
		It("Should return a queue", func() {
			for i:=1;i<=5;i++{
				testarrayqueue.Enqueue(i)
			}
			testarrayqueue.Dequeue()
			testvalue := make([]interface{},5,5)
			for i:=1;i<=4;i++{
				testvalue[i] = i+1
			}
			Expect(testarrayqueue.Candequeue()).To(Equal(true))
			Expect(testarrayqueue.Data).To(Equal(testvalue))
		})
	})
	Describe("Now Let dequeue and it should move the data as expected",func(){
		It("Should move data as expected", func() {
			for i:=1;i<=5;i++{
				testarrayqueue.Enqueue(i)
			}
			testarrayqueue.Dequeue()
			testarrayqueue.Dequeue()

			testarrayqueue.Enqueue(6)
			testarrayqueue.Enqueue(7)

			testvalue := make([]interface{},5,5)
			for i:=0;i<5;i++{
				testvalue[i] = i+3
			}
			Expect(testarrayqueue.Candequeue()).To(Equal(true))
			Expect(testarrayqueue.Data).To(Equal(testvalue))
		})
	})


})
