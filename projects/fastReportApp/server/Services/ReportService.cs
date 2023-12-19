using Grpc.Core;
using FastReportApp;

namespace Server.Services;

public class ReportService : report.reportBase
{
    private readonly ILogger<ReportService> _logger;
    public ReportService(ILogger<ReportService> logger)
    {
        _logger = logger;
    }

    public override Task<GetReportResponse> GetReport(GetReportRequest request, ServerCallContext context)
    {
        Console.WriteLine("Olá, mundo!");
        GetReportResponse response = new GetReportResponse();
        // Lógica para preencher a resposta, se necessário
        
        return Task.FromResult(response);
    }
}
