package cyclequeue_test

import (
	"Algorithm-Datastruct/golang/Queue/cyclequeue"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cyclequeue", func() {

	var testCyclequeue *cyclequeue.CycleQueue
	BeforeEach(
		func(){testCyclequeue = cyclequeue.NewCycleQueue(5)},
		)
	Context("should return a cyclequeue", func() {
		It("should get a cyclenque", func() {
			Expect(testCyclequeue.Data).To(Equal(make([]interface{},5,5)))
		})
	})
	Context("test the cyclequeue", func() {
		It("after do enqueue, it should return a good queue", func(){
			for i:=1;i<=5;i++{
				testCyclequeue.Enqueue(i)
			}
			expectedresult:=make([]interface{},5,5)
			//会浪费一个空间的
			for i:=0;i<4;i++{
				expectedresult[i]=i+1
			}
			Expect(testCyclequeue.Data).To(Equal(expectedresult))
		})
	})

	Context("test the cyclequeue", func() {
		It("after do enqueue and deenqueue, it should return a good queue", func(){
			for i:=1;i<=5;i++{
				testCyclequeue.Enqueue(i)
			}
			//1 2 3 4 nil
			testCyclequeue.Dequeue()
			testCyclequeue.Dequeue()
			//nil nil 3 4 nil
			expectedresult:=make([]interface{},5,5)
			//会浪费一个空间的
			for i:=2;i<4;i++{
				expectedresult[i]=i+1
				//ex:nil nil 4 5 nil
			}
			Expect(testCyclequeue.Data).To(Equal(expectedresult))


			testCyclequeue.Enqueue(6)
			//nil nil 4 5 6
			expectedresult[4] = 6
			//ex nil nil 4 5 6
			Expect(testCyclequeue.Data).To(Equal(expectedresult))

			testCyclequeue.Enqueue(7)
			//7 nil 4 5 6
			expectedresult[0] = 7
			//ex 7 nil 4 5 6
			Expect(testCyclequeue.Data).To(Equal(expectedresult))

			testCyclequeue.Enqueue(8)
			//7 nil 4 5 6
			expectedresult[1] = nil
			//ex 7 nil 4 5 6
			Expect(testCyclequeue.Data).To(Equal(expectedresult))

			testCyclequeue.Dequeue()
			//7 nil nil 5 6
			expectedresult[2] = nil
			Expect(testCyclequeue.Data).To(Equal(expectedresult))

		})
	})




})
