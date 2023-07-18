package gocha

import "testing"

var pieChartConfig string = `
{
	type: 'pie',
	data: {
	  datasets: [
		{
		  data: [84, 28, 57, 19, 97],
		  backgroundColor: [
			'rgb(255, 99, 132)',
			'rgb(255, 159, 64)',
			'rgb(255, 205, 86)',
			'rgb(75, 192, 192)',
			'rgb(54, 162, 235)',
		  ],
		  label: 'Dataset 1',
		},
	  ],
	  labels: ['Red', 'Orange', 'Yellow', 'Green', 'Blue'],
	},
  }
`

var pieChartURL string = "https://quickchart.io/chart?bkg=&c=%0A%7B%0A%09type%3A+%27pie%27%2C%0A%09data%3A+%7B%0A%09++datasets%3A+%5B%0A%09%09%7B%0A%09%09++data%3A+%5B84%2C+28%2C+57%2C+19%2C+97%5D%2C%0A%09%09++backgroundColor%3A+%5B%0A%09%09%09%27rgb%28255%2C+99%2C+132%29%27%2C%0A%09%09%09%27rgb%28255%2C+159%2C+64%29%27%2C%0A%09%09%09%27rgb%28255%2C+205%2C+86%29%27%2C%0A%09%09%09%27rgb%2875%2C+192%2C+192%29%27%2C%0A%09%09%09%27rgb%2854%2C+162%2C+235%29%27%2C%0A%09%09++%5D%2C%0A%09%09++label%3A+%27Dataset+1%27%2C%0A%09%09%7D%2C%0A%09++%5D%2C%0A%09++labels%3A+%5B%27Red%27%2C+%27Orange%27%2C+%27Yellow%27%2C+%27Green%27%2C+%27Blue%27%5D%2C%0A%09%7D%2C%0A++%7D%0A&devicePixelRatio=&f=&h=&v=&w="

func TestQuickChart_GetURL(t *testing.T) {
	type fields struct {
		config           string
		width            string
		height           string
		devicePixelRatio string
		backgroundColor  string
		version          string
		format           chartFormat
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"Without config", fields{}, "", true},
		{"Pie Chart config", fields{config: pieChartConfig}, pieChartURL, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qc := &QuickChart{
				config:           tt.fields.config,
				width:            tt.fields.width,
				height:           tt.fields.height,
				devicePixelRatio: tt.fields.devicePixelRatio,
				backgroundColor:  tt.fields.backgroundColor,
				version:          tt.fields.version,
				format:           tt.fields.format,
			}
			got, err := qc.GetURL()
			if (err != nil) != tt.wantErr {
				t.Errorf("QuickChart.GetURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QuickChart.GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
