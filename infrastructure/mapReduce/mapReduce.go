package mapReduce

// Flow:
// 1) split input data to channels
// 2) input function => sent split data into reducer
// 3) divide data into diapasons of different periods
// 4) map function => apply any function to each sub data
// 5) reduce function => join results of reducers into 1 collection

// MapperCollector is a channel that collects the output from mapper tasks
type MapperCollector chan chan interface{}

// MapperFunc is a function that performs the mapping part of the MapReduce job
type MapperFunc func(interface{}, chan interface{})

// ReducerFunc is a function that performs the reduce part of the MapReduce job
type ReducerFunc func(chan interface{}, chan interface{})

// The mapperDispatcher function is responsible to listen on the data channel that receives each filename
// to be processed and invoke a mapper for each file, pushing the output of the job into a MapperCollector
func mapperDispatcher(mapper MapperFunc, input chan interface{}, collector MapperCollector) {
	for item := range input {
		taskOutput := make(chan interface{})
		go mapper(item, taskOutput)
		collector <- taskOutput
	}
	close(collector)
}

// The reducerDispatcher function is responsible to listen on the collector channel
// and push each item as the data for the Reducer task.
func reducerDispatcher(collector MapperCollector, reducerInput chan interface{}) {
	for output := range collector {
		reducerInput <- <-output
	}
	close(reducerInput)
}

const (
	MaxWorkers = 10
)

func MapReduce(mapper MapperFunc, reducer ReducerFunc, input chan interface{}) interface{} {

	reducerInput := make(chan interface{})
	reducerOutput := make(chan interface{})
	mapperCollector := make(MapperCollector, MaxWorkers)

	go reducer(reducerInput, reducerOutput)
	go reducerDispatcher(mapperCollector, reducerInput)
	go mapperDispatcher(mapper, input, mapperCollector)
	return <-reducerOutput

}