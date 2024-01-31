void main() {
  List<double> faturamentoUltimosMeses = [
    5397.28,
    0,
    0,
    25760.62,
    23464.03,
    0,
    10464.07,
    22978.16,
    21911.17,
    7465.06,
    24781.98,
    23452.10,
    26183.04,
  ];

  List<double> folhaPagamentoUltimosMeses = [
    0,
    0,
    0,
    26400,
    6570,
    6570,
    6570,
    4334 + 2100,
    4036 + 2100,
    6570 + 2100,
    3470 + 3470,
    3284 + 3284,
  ];

  double totalFaturadoUltimos12Meses =
  getTotalUltimos12Meses(faturamentoUltimosMeses);
  double totalFolha12Meses = getTotalUltimos12Meses(folhaPagamentoUltimosMeses);
  double fatorR = totalFolha12Meses / totalFaturadoUltimos12Meses * 100;

  print(
      "Total faturado últimos 12 meses: ${totalFaturadoUltimos12Meses.toStringAsFixed(2)}, folha de pagamento últimos 12 meses: ${totalFolha12Meses.toStringAsFixed(2)}, fator R: ${fatorR.toStringAsFixed(2)}%");
}

double getTotalUltimos12Meses(List<double> lista) {
  if (lista.length > 12) {
    lista = lista.sublist(lista.length - 12);
  }
  double total = 0;
  for (double v in lista) {
    total += v;
  }
  return total;
}
