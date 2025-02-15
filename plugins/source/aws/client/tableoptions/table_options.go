package tableoptions

type TableOptions struct {
	CloudwatchMetrics []CloudwatchMetrics `json:"aws_alpha_cloudwatch_metrics,omitempty"`

	CloudTrailEvents       *CloudtrailAPIs         `json:"aws_cloudtrail_events,omitempty"`
	AccessAnalyzerFindings *AccessanalyzerFindings `json:"aws_accessanalyzer_analyzer_findings,omitempty"`
	Inspector2Findings     *Inspector2APIs         `json:"aws_inspector2_findings,omitempty"`
}

func (t TableOptions) Validate() error {
	if t.CloudTrailEvents != nil {
		err := t.CloudTrailEvents.Validate()
		if err != nil {
			return err
		}
	}

	if t.AccessAnalyzerFindings != nil {
		err := t.AccessAnalyzerFindings.Validate()
		if err != nil {
			return err
		}
	}

	if t.Inspector2Findings != nil {
		err := t.Inspector2Findings.Validate()
		if err != nil {
			return err
		}
	}

	for _, m := range t.CloudwatchMetrics {
		if err := m.Validate(); err != nil {
			return err
		}
	}

	return nil
}
