var reporter = require('cucumber-html-reporter');

var options = {
    theme: 'bootstrap',
    jsonFile: 'log/report.json',
    output: 'log/report.html',
    reportSuiteAsScenarios: true,
    launchReport: true,
    metadata: {
        "App Version":"0.1.1",
        "Test Environment": "test",
        "Browser": "Chrome",
        "Platform": "MacOS",
        "Parallel": "Scenarios",
        "Executed": "Remote"
    }
};

reporter.generate(options);