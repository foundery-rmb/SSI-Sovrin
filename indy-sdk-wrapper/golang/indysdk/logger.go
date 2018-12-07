package indysdk

import "github.com/go-kit/kit/log"

logger := log.NewLogfmtLogger(os.Stdout)

// make time predictable for this test
baseTime := time.Date(2015, time.February, 3, 10, 0, 0, 0, time.UTC)
mockTime := func() time.Time {
    baseTime = baseTime.Add(time.Second)
    return baseTime
}

logger = log.With(logger, "time", log.Timestamp(mockTime), "caller", log.DefaultCaller)