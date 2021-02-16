package bbolt

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	mmapSize = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bbolt_mmap_size",
		Help: "Size of the mmap allocation in bytes",
	})
	numMmapResizes = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bbolt_mmap_resize_count_total",
		Help: "Total number of times we have resized mmap",
	})
	readOnlyTxCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bbolt_read_only_tx_total",
		Help: "Total number of times we started read only txs",
	})
	writableTxCount = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bbolt_read_write_tx_total",
		Help: "Total number of times we started writeable txs",
	})
	readWriteTxLockWaitTime = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "read_write_tx_lock_time_milliseconds",
			Help:    "The number of time waiting for the db read/write lock in milliseconds",
			Buckets: []float64{100, 500, 1000, 2000, 4000, 10000},
		},
	)
	readTxLockWaitTime = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "read_tx_lock_time_milliseconds",
			Help:    "The number of time waiting for the db read lock in milliseconds",
			Buckets: []float64{100, 500, 1000, 2000, 4000, 10000},
		},
	)
	mmapWaitTime = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "mmap_time_milliseconds",
			Help:    "The number of time waiting for mmap to complete in milliseconds",
			Buckets: []float64{100, 500, 1000, 2000, 4000, 10000, 20000},
		},
	)
)
