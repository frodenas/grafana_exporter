package test_matchers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/prometheus/client_golang/prometheus"

	. "github.com/frodenas/grafana_exporter/utils/test_matchers"
)

var _ = Describe("PrometheusMetric", func() {
	var (
		metricNamespace       = "fake_namespace"
		metricSubsystem       = "fake_sybsystem"
		metricName            = "fake_name"
		metricHelp            = "Fake Metric Help"
		metricLabelName       = "fake_label_name"
		metricLabelValue      = "fake_label_value"
		metricConstLabelName  = "fake_constant_label_name"
		metricConstLabelValue = "fake_constant_label_value"
	)

	Context("When asserting equality between Counter Metrics", func() {
		It("should do the right thing", func() {
			expectedMetric := prometheus.NewCounter(
				prometheus.CounterOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				})
			expectedMetric.Inc()

			actualMetric := prometheus.NewCounter(
				prometheus.CounterOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				})
			actualMetric.Inc()

			Expect(expectedMetric).To(PrometheusMetric(actualMetric))
		})
	})

	Context("When asserting equality between CounterVec Metrics", func() {
		It("should do the right thing", func() {
			expectedMetric := prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				},
				[]string{metricLabelName},
			)
			expectedMetric.WithLabelValues(metricLabelValue).Inc()

			actualMetric := prometheus.NewCounterVec(
				prometheus.CounterOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				},
				[]string{metricLabelName},
			)
			actualMetric.WithLabelValues(metricLabelValue).Inc()

			Expect(expectedMetric.WithLabelValues(metricLabelValue)).To(PrometheusMetric(actualMetric.WithLabelValues(metricLabelValue)))
		})
	})

	Context("When asserting equality between Gauge Metrics", func() {
		It("should do the right thing", func() {
			expectedMetric := prometheus.NewGauge(
				prometheus.GaugeOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				})
			expectedMetric.Inc()

			actualMetric := prometheus.NewGauge(
				prometheus.GaugeOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				})
			actualMetric.Inc()

			Expect(expectedMetric).To(PrometheusMetric(actualMetric))
		})
	})

	Context("When asserting equality between GaugeVec Metrics", func() {
		It("should do the right thing", func() {
			expectedMetric := prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				},
				[]string{metricLabelName},
			)
			expectedMetric.WithLabelValues(metricLabelValue).Inc()

			actualMetric := prometheus.NewGaugeVec(
				prometheus.GaugeOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				},
				[]string{metricLabelName},
			)
			actualMetric.WithLabelValues(metricLabelValue).Inc()

			Expect(expectedMetric.WithLabelValues(metricLabelValue)).To(PrometheusMetric(actualMetric.WithLabelValues(metricLabelValue)))
		})
	})

	Context("When asserting equality between Histogram Metrics", func() {
		It("should do the right thing", func() {
			expectedMetric := prometheus.NewHistogram(
				prometheus.HistogramOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				})
			expectedMetric.Observe(float64(1))

			actualMetric := prometheus.NewHistogram(
				prometheus.HistogramOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				})
			actualMetric.Observe(float64(1))

			Expect(expectedMetric).To(PrometheusMetric(actualMetric))
		})
	})

	Context("When asserting equality between Summary Metrics", func() {
		It("should do the right thing", func() {
			expectedMetric := prometheus.NewSummary(
				prometheus.SummaryOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				})
			expectedMetric.Observe(float64(1))

			actualMetric := prometheus.NewSummary(
				prometheus.SummaryOpts{
					Namespace:   metricNamespace,
					Subsystem:   metricSubsystem,
					Name:        metricName,
					Help:        metricHelp,
					ConstLabels: prometheus.Labels{metricConstLabelName: metricConstLabelValue},
				})
			actualMetric.Observe(float64(1))

			Expect(expectedMetric).To(PrometheusMetric(actualMetric))
		})
	})
})
