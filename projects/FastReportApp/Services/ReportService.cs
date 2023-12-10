using Grpc.Core;
using FastReportApp;

namespace FastReportApp.Services;

public class ReportService : report.reportBase
{
    private readonly ILogger<ReportService> _logger;
    public ReportService(ILogger<ReportService> logger)
    {
        _logger = logger;
    }
    
    public override Task<GetReportResponse> GetReport(GetReportRequest request, ServerCallContext context)
    {
        return Task.FromResult(new GetReportResponse
        {
           Data = "" // report content
        });
    }
}
