import * as jspb from 'google-protobuf'

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb'; // proto import: "google/protobuf/timestamp.proto"
import * as google_api_annotations_pb from '../../../google/api/annotations_pb'; // proto import: "google/api/annotations.proto"
import * as google_api_resource_pb from '../../../google/api/resource_pb'; // proto import: "google/api/resource.proto"
import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb'; // proto import: "google/protobuf/field_mask.proto"
import * as plugins_service_service_pb from '../../../plugins/service/service_pb'; // proto import: "plugins/service/service.proto"
import * as plugins_validate_validate_pb from '../../../plugins/validate/validate_pb'; // proto import: "plugins/validate/validate.proto"
import * as filter_filter_pb from '../../../filter/filter_pb'; // proto import: "filter/filter.proto"


export class CreateOrderRequest extends jspb.Message {
  getOrder(): Order | undefined;
  setOrder(value?: Order): CreateOrderRequest;
  hasOrder(): boolean;
  clearOrder(): CreateOrderRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateOrderRequest): CreateOrderRequest.AsObject;
  static serializeBinaryToWriter(message: CreateOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateOrderRequest;
  static deserializeBinaryFromReader(message: CreateOrderRequest, reader: jspb.BinaryReader): CreateOrderRequest;
}

export namespace CreateOrderRequest {
  export type AsObject = {
    order?: Order.AsObject,
  }
}

export class CreateOrderResponse extends jspb.Message {
  getOrder(): Order | undefined;
  setOrder(value?: Order): CreateOrderResponse;
  hasOrder(): boolean;
  clearOrder(): CreateOrderResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateOrderResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CreateOrderResponse): CreateOrderResponse.AsObject;
  static serializeBinaryToWriter(message: CreateOrderResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateOrderResponse;
  static deserializeBinaryFromReader(message: CreateOrderResponse, reader: jspb.BinaryReader): CreateOrderResponse;
}

export namespace CreateOrderResponse {
  export type AsObject = {
    order?: Order.AsObject,
  }
}

export class UpdateOrderRequest extends jspb.Message {
  getId(): string;
  setId(value: string): UpdateOrderRequest;

  getOrder(): Order | undefined;
  setOrder(value?: Order): UpdateOrderRequest;
  hasOrder(): boolean;
  clearOrder(): UpdateOrderRequest;

  getUpdateMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setUpdateMask(value?: google_protobuf_field_mask_pb.FieldMask): UpdateOrderRequest;
  hasUpdateMask(): boolean;
  clearUpdateMask(): UpdateOrderRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateOrderRequest): UpdateOrderRequest.AsObject;
  static serializeBinaryToWriter(message: UpdateOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateOrderRequest;
  static deserializeBinaryFromReader(message: UpdateOrderRequest, reader: jspb.BinaryReader): UpdateOrderRequest;
}

export namespace UpdateOrderRequest {
  export type AsObject = {
    id: string,
    order?: Order.AsObject,
    updateMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
  }
}

export class UpdateOrderResponse extends jspb.Message {
  getOrder(): Order | undefined;
  setOrder(value?: Order): UpdateOrderResponse;
  hasOrder(): boolean;
  clearOrder(): UpdateOrderResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UpdateOrderResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UpdateOrderResponse): UpdateOrderResponse.AsObject;
  static serializeBinaryToWriter(message: UpdateOrderResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UpdateOrderResponse;
  static deserializeBinaryFromReader(message: UpdateOrderResponse, reader: jspb.BinaryReader): UpdateOrderResponse;
}

export namespace UpdateOrderResponse {
  export type AsObject = {
    order?: Order.AsObject,
  }
}

export class DeleteOrderRequest extends jspb.Message {
  getId(): string;
  setId(value: string): DeleteOrderRequest;

  getHard(): boolean;
  setHard(value: boolean): DeleteOrderRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteOrderRequest): DeleteOrderRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteOrderRequest;
  static deserializeBinaryFromReader(message: DeleteOrderRequest, reader: jspb.BinaryReader): DeleteOrderRequest;
}

export namespace DeleteOrderRequest {
  export type AsObject = {
    id: string,
    hard: boolean,
  }
}

export class DeleteOrderResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteOrderResponse.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteOrderResponse): DeleteOrderResponse.AsObject;
  static serializeBinaryToWriter(message: DeleteOrderResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteOrderResponse;
  static deserializeBinaryFromReader(message: DeleteOrderResponse, reader: jspb.BinaryReader): DeleteOrderResponse;
}

export namespace DeleteOrderResponse {
  export type AsObject = {
  }
}

export class ListOrderRequest extends jspb.Message {
  getId(): string;
  setId(value: string): ListOrderRequest;

  getNome(): string;
  setNome(value: string): ListOrderRequest;

  getFilter(): filter_filter_pb.Filter | undefined;
  setFilter(value?: filter_filter_pb.Filter): ListOrderRequest;
  hasFilter(): boolean;
  clearFilter(): ListOrderRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListOrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListOrderRequest): ListOrderRequest.AsObject;
  static serializeBinaryToWriter(message: ListOrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListOrderRequest;
  static deserializeBinaryFromReader(message: ListOrderRequest, reader: jspb.BinaryReader): ListOrderRequest;
}

export namespace ListOrderRequest {
  export type AsObject = {
    id: string,
    nome: string,
    filter?: filter_filter_pb.Filter.AsObject,
  }
}

export class ListOrderResponse extends jspb.Message {
  getOrdersList(): Array<Order>;
  setOrdersList(value: Array<Order>): ListOrderResponse;
  clearOrdersList(): ListOrderResponse;
  addOrders(value?: Order, index?: number): Order;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListOrderResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListOrderResponse): ListOrderResponse.AsObject;
  static serializeBinaryToWriter(message: ListOrderResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListOrderResponse;
  static deserializeBinaryFromReader(message: ListOrderResponse, reader: jspb.BinaryReader): ListOrderResponse;
}

export namespace ListOrderResponse {
  export type AsObject = {
    ordersList: Array<Order.AsObject>,
  }
}

export class CheckoutRequest extends jspb.Message {
  getLinkid(): string;
  setLinkid(value: string): CheckoutRequest;

  getOrderid(): string;
  setOrderid(value: string): CheckoutRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CheckoutRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CheckoutRequest): CheckoutRequest.AsObject;
  static serializeBinaryToWriter(message: CheckoutRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CheckoutRequest;
  static deserializeBinaryFromReader(message: CheckoutRequest, reader: jspb.BinaryReader): CheckoutRequest;
}

export namespace CheckoutRequest {
  export type AsObject = {
    linkid: string,
    orderid: string,
  }
}

export class CheckoutResponse extends jspb.Message {
  getOrder(): Order | undefined;
  setOrder(value?: Order): CheckoutResponse;
  hasOrder(): boolean;
  clearOrder(): CheckoutResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CheckoutResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CheckoutResponse): CheckoutResponse.AsObject;
  static serializeBinaryToWriter(message: CheckoutResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CheckoutResponse;
  static deserializeBinaryFromReader(message: CheckoutResponse, reader: jspb.BinaryReader): CheckoutResponse;
}

export namespace CheckoutResponse {
  export type AsObject = {
    order?: Order.AsObject,
  }
}

export class Order extends jspb.Message {
  getId(): string;
  setId(value: string): Order;

  getOrigem(): Origem;
  setOrigem(value: Origem): Order;

  getTipo(): TipoMovimentacao;
  setTipo(value: TipoMovimentacao): Order;

  getSituacao(): TipoSituacao;
  setSituacao(value: TipoSituacao): Order;

  getTipoMoeda(): TipoMoeda;
  setTipoMoeda(value: TipoMoeda): Order;

  getTabelaPreco(): TabelaPreco;
  setTabelaPreco(value: TabelaPreco): Order;

  getRecorrenciaCriada(): boolean;
  setRecorrenciaCriada(value: boolean): Order;

  getPeriodicidadeTipo(): Periodicidade;
  setPeriodicidadeTipo(value: Periodicidade): Order;

  getPeriodicidadeValor(): number;
  setPeriodicidadeValor(value: number): Order;

  getPeriodicidadeData(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setPeriodicidadeData(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasPeriodicidadeData(): boolean;
  clearPeriodicidadeData(): Order;

  getImportadoEm(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setImportadoEm(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasImportadoEm(): boolean;
  clearImportadoEm(): Order;

  getCobrancaEnviar(): boolean;
  setCobrancaEnviar(value: boolean): Order;

  getCobrancaDataUltima(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCobrancaDataUltima(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasCobrancaDataUltima(): boolean;
  clearCobrancaDataUltima(): Order;

  getPessoa(): Pessoa | undefined;
  setPessoa(value?: Pessoa): Order;
  hasPessoa(): boolean;
  clearPessoa(): Order;

  getVendedor(): Pessoa | undefined;
  setVendedor(value?: Pessoa): Order;
  hasVendedor(): boolean;
  clearVendedor(): Order;

  getNumero(): number;
  setNumero(value: number): Order;

  getDataHoraRegistro(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDataHoraRegistro(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasDataHoraRegistro(): boolean;
  clearDataHoraRegistro(): Order;

  getDataHoraFechamento(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDataHoraFechamento(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasDataHoraFechamento(): boolean;
  clearDataHoraFechamento(): Order;

  getPrevisaoEntrega(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setPrevisaoEntrega(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasPrevisaoEntrega(): boolean;
  clearPrevisaoEntrega(): Order;

  getPrevisaoEntregaDescricao(): string;
  setPrevisaoEntregaDescricao(value: string): Order;

  getDataHoraInicio(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDataHoraInicio(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasDataHoraInicio(): boolean;
  clearDataHoraInicio(): Order;

  getDataHoraEntrega(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDataHoraEntrega(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasDataHoraEntrega(): boolean;
  clearDataHoraEntrega(): Order;

  getTempoDecorrido(): string;
  setTempoDecorrido(value: string): Order;

  getDescontoValor(): number;
  setDescontoValor(value: number): Order;

  getAcrescimoValor(): number;
  setAcrescimoValor(value: number): Order;

  getValorSubtotal(): number;
  setValorSubtotal(value: number): Order;

  getValorTotalServico(): number;
  setValorTotalServico(value: number): Order;

  getValorTotalProduto(): number;
  setValorTotalProduto(value: number): Order;

  getValorTotal(): number;
  setValorTotal(value: number): Order;

  getValorTotalPago(): number;
  setValorTotalPago(value: number): Order;

  getCashbackValor(): number;
  setCashbackValor(value: number): Order;

  getComissaoValor(): number;
  setComissaoValor(value: number): Order;

  getDfe(): Dfe | undefined;
  setDfe(value?: Dfe): Order;
  hasDfe(): boolean;
  clearDfe(): Order;

  getObs(): string;
  setObs(value: string): Order;

  getCondicaoPagamentoNome(): string;
  setCondicaoPagamentoNome(value: string): Order;

  getCondicaoPagamentoId(): string;
  setCondicaoPagamentoId(value: string): Order;

  getCancelamentoMotivo(): string;
  setCancelamentoMotivo(value: string): Order;

  getCancelamentoUsuarioId(): string;
  setCancelamentoUsuarioId(value: string): Order;

  getCancelamentoUsuarioNome(): string;
  setCancelamentoUsuarioNome(value: string): Order;

  getCancelamentoDataHora(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCancelamentoDataHora(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasCancelamentoDataHora(): boolean;
  clearCancelamentoDataHora(): Order;

  getProdutosList(): Array<Produto>;
  setProdutosList(value: Array<Produto>): Order;
  clearProdutosList(): Order;
  addProdutos(value?: Produto, index?: number): Produto;

  getServicosList(): Array<Servico>;
  setServicosList(value: Array<Servico>): Order;
  clearServicosList(): Order;
  addServicos(value?: Servico, index?: number): Servico;

  getFormasPagamentoList(): Array<FormaPagamento>;
  setFormasPagamentoList(value: Array<FormaPagamento>): Order;
  clearFormasPagamentoList(): Order;
  addFormasPagamento(value?: FormaPagamento, index?: number): FormaPagamento;

  getDataEntregaServico(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDataEntregaServico(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasDataEntregaServico(): boolean;
  clearDataEntregaServico(): Order;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Order.AsObject;
  static toObject(includeInstance: boolean, msg: Order): Order.AsObject;
  static serializeBinaryToWriter(message: Order, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Order;
  static deserializeBinaryFromReader(message: Order, reader: jspb.BinaryReader): Order;
}

export namespace Order {
  export type AsObject = {
    id: string,
    origem: Origem,
    tipo: TipoMovimentacao,
    situacao: TipoSituacao,
    tipoMoeda: TipoMoeda,
    tabelaPreco: TabelaPreco,
    recorrenciaCriada: boolean,
    periodicidadeTipo: Periodicidade,
    periodicidadeValor: number,
    periodicidadeData?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    importadoEm?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    cobrancaEnviar: boolean,
    cobrancaDataUltima?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    pessoa?: Pessoa.AsObject,
    vendedor?: Pessoa.AsObject,
    numero: number,
    dataHoraRegistro?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    dataHoraFechamento?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    previsaoEntrega?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    previsaoEntregaDescricao: string,
    dataHoraInicio?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    dataHoraEntrega?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    tempoDecorrido: string,
    descontoValor: number,
    acrescimoValor: number,
    valorSubtotal: number,
    valorTotalServico: number,
    valorTotalProduto: number,
    valorTotal: number,
    valorTotalPago: number,
    cashbackValor: number,
    comissaoValor: number,
    dfe?: Dfe.AsObject,
    obs: string,
    condicaoPagamentoNome: string,
    condicaoPagamentoId: string,
    cancelamentoMotivo: string,
    cancelamentoUsuarioId: string,
    cancelamentoUsuarioNome: string,
    cancelamentoDataHora?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    produtosList: Array<Produto.AsObject>,
    servicosList: Array<Servico.AsObject>,
    formasPagamentoList: Array<FormaPagamento.AsObject>,
    dataEntregaServico?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Produto extends jspb.Message {
  getId(): string;
  setId(value: string): Produto;

  getNome(): string;
  setNome(value: string): Produto;

  getUnidadeMedida(): string;
  setUnidadeMedida(value: string): Produto;

  getQuantidade(): number;
  setQuantidade(value: number): Produto;

  getValorUnitario(): number;
  setValorUnitario(value: number): Produto;

  getValorTotal(): number;
  setValorTotal(value: number): Produto;

  getValorDesconto(): number;
  setValorDesconto(value: number): Produto;

  getValorAcrescimo(): number;
  setValorAcrescimo(value: number): Produto;

  getValorTotalLiquido(): number;
  setValorTotalLiquido(value: number): Produto;

  getCancelado(): boolean;
  setCancelado(value: boolean): Produto;

  getObservacao(): string;
  setObservacao(value: string): Produto;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Produto.AsObject;
  static toObject(includeInstance: boolean, msg: Produto): Produto.AsObject;
  static serializeBinaryToWriter(message: Produto, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Produto;
  static deserializeBinaryFromReader(message: Produto, reader: jspb.BinaryReader): Produto;
}

export namespace Produto {
  export type AsObject = {
    id: string,
    nome: string,
    unidadeMedida: string,
    quantidade: number,
    valorUnitario: number,
    valorTotal: number,
    valorDesconto: number,
    valorAcrescimo: number,
    valorTotalLiquido: number,
    cancelado: boolean,
    observacao: string,
  }
}

export class Servico extends jspb.Message {
  getId(): string;
  setId(value: string): Servico;

  getNome(): string;
  setNome(value: string): Servico;

  getUnidadeMedida(): string;
  setUnidadeMedida(value: string): Servico;

  getQuantidade(): number;
  setQuantidade(value: number): Servico;

  getValorUnitario(): number;
  setValorUnitario(value: number): Servico;

  getValorTotal(): number;
  setValorTotal(value: number): Servico;

  getValorDesconto(): number;
  setValorDesconto(value: number): Servico;

  getValorAcrescimo(): number;
  setValorAcrescimo(value: number): Servico;

  getValorTotalLiquido(): number;
  setValorTotalLiquido(value: number): Servico;

  getCancelado(): boolean;
  setCancelado(value: boolean): Servico;

  getObservacao(): string;
  setObservacao(value: string): Servico;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Servico.AsObject;
  static toObject(includeInstance: boolean, msg: Servico): Servico.AsObject;
  static serializeBinaryToWriter(message: Servico, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Servico;
  static deserializeBinaryFromReader(message: Servico, reader: jspb.BinaryReader): Servico;
}

export namespace Servico {
  export type AsObject = {
    id: string,
    nome: string,
    unidadeMedida: string,
    quantidade: number,
    valorUnitario: number,
    valorTotal: number,
    valorDesconto: number,
    valorAcrescimo: number,
    valorTotalLiquido: number,
    cancelado: boolean,
    observacao: string,
  }
}

export class FormaPagamento extends jspb.Message {
  getId(): string;
  setId(value: string): FormaPagamento;

  getNome(): string;
  setNome(value: string): FormaPagamento;

  getValor(): number;
  setValor(value: number): FormaPagamento;

  getNumeroAutorizacao(): string;
  setNumeroAutorizacao(value: string): FormaPagamento;

  getSituacao(): SituacaoPagamento;
  setSituacao(value: SituacaoPagamento): FormaPagamento;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FormaPagamento.AsObject;
  static toObject(includeInstance: boolean, msg: FormaPagamento): FormaPagamento.AsObject;
  static serializeBinaryToWriter(message: FormaPagamento, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FormaPagamento;
  static deserializeBinaryFromReader(message: FormaPagamento, reader: jspb.BinaryReader): FormaPagamento;
}

export namespace FormaPagamento {
  export type AsObject = {
    id: string,
    nome: string,
    valor: number,
    numeroAutorizacao: string,
    situacao: SituacaoPagamento,
  }
}

export class Pessoa extends jspb.Message {
  getId(): string;
  setId(value: string): Pessoa;

  getRevenda(): boolean;
  setRevenda(value: boolean): Pessoa;

  getCpfCnpj(): string;
  setCpfCnpj(value: string): Pessoa;

  getIe(): string;
  setIe(value: string): Pessoa;

  getNome(): string;
  setNome(value: string): Pessoa;

  getNome2(): string;
  setNome2(value: string): Pessoa;

  getUsuarioId(): string;
  setUsuarioId(value: string): Pessoa;

  getUsuarioNome(): string;
  setUsuarioNome(value: string): Pessoa;

  getComissao(): number;
  setComissao(value: number): Pessoa;

  getEndNome(): string;
  setEndNome(value: string): Pessoa;

  getEndCep(): string;
  setEndCep(value: string): Pessoa;

  getEndEndereco(): string;
  setEndEndereco(value: string): Pessoa;

  getEndNumero(): string;
  setEndNumero(value: string): Pessoa;

  getEndBairro(): string;
  setEndBairro(value: string): Pessoa;

  getEndCidade(): string;
  setEndCidade(value: string): Pessoa;

  getEndCidadeCodigo(): string;
  setEndCidadeCodigo(value: string): Pessoa;

  getEndUf(): string;
  setEndUf(value: string): Pessoa;

  getTelefone(): string;
  setTelefone(value: string): Pessoa;

  getEmail(): string;
  setEmail(value: string): Pessoa;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Pessoa.AsObject;
  static toObject(includeInstance: boolean, msg: Pessoa): Pessoa.AsObject;
  static serializeBinaryToWriter(message: Pessoa, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Pessoa;
  static deserializeBinaryFromReader(message: Pessoa, reader: jspb.BinaryReader): Pessoa;
}

export namespace Pessoa {
  export type AsObject = {
    id: string,
    revenda: boolean,
    cpfCnpj: string,
    ie: string,
    nome: string,
    nome2: string,
    usuarioId: string,
    usuarioNome: string,
    comissao: number,
    endNome: string,
    endCep: string,
    endEndereco: string,
    endNumero: string,
    endBairro: string,
    endCidade: string,
    endCidadeCodigo: string,
    endUf: string,
    telefone: string,
    email: string,
  }
}

export class Dfe extends jspb.Message {
  getNfe(): Nfe | undefined;
  setNfe(value?: Nfe): Dfe;
  hasNfe(): boolean;
  clearNfe(): Dfe;

  getNfce(): Nfce | undefined;
  setNfce(value?: Nfce): Dfe;
  hasNfce(): boolean;
  clearNfce(): Dfe;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Dfe.AsObject;
  static toObject(includeInstance: boolean, msg: Dfe): Dfe.AsObject;
  static serializeBinaryToWriter(message: Dfe, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Dfe;
  static deserializeBinaryFromReader(message: Dfe, reader: jspb.BinaryReader): Dfe;
}

export namespace Dfe {
  export type AsObject = {
    nfe?: Nfe.AsObject,
    nfce?: Nfce.AsObject,
  }
}

export class Nfe extends jspb.Message {
  getId(): string;
  setId(value: string): Nfe;

  getTipo(): string;
  setTipo(value: string): Nfe;

  getSituacao(): string;
  setSituacao(value: string): Nfe;

  getChave(): string;
  setChave(value: string): Nfe;

  getUrlDanfe(): string;
  setUrlDanfe(value: string): Nfe;

  getSerie(): number;
  setSerie(value: number): Nfe;

  getNumero(): number;
  setNumero(value: number): Nfe;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Nfe.AsObject;
  static toObject(includeInstance: boolean, msg: Nfe): Nfe.AsObject;
  static serializeBinaryToWriter(message: Nfe, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Nfe;
  static deserializeBinaryFromReader(message: Nfe, reader: jspb.BinaryReader): Nfe;
}

export namespace Nfe {
  export type AsObject = {
    id: string,
    tipo: string,
    situacao: string,
    chave: string,
    urlDanfe: string,
    serie: number,
    numero: number,
  }
}

export class Nfce extends jspb.Message {
  getNfceNumero(): number;
  setNfceNumero(value: number): Nfce;

  getNfceUrlDanfe(): string;
  setNfceUrlDanfe(value: string): Nfce;

  getNfceUrlXml(): string;
  setNfceUrlXml(value: string): Nfce;

  getNfceSerie(): number;
  setNfceSerie(value: number): Nfce;

  getNfceChave(): string;
  setNfceChave(value: string): Nfce;

  getNfceId(): string;
  setNfceId(value: string): Nfce;

  getNfceSituacao(): string;
  setNfceSituacao(value: string): Nfce;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Nfce.AsObject;
  static toObject(includeInstance: boolean, msg: Nfce): Nfce.AsObject;
  static serializeBinaryToWriter(message: Nfce, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Nfce;
  static deserializeBinaryFromReader(message: Nfce, reader: jspb.BinaryReader): Nfce;
}

export namespace Nfce {
  export type AsObject = {
    nfceNumero: number,
    nfceUrlDanfe: string,
    nfceUrlXml: string,
    nfceSerie: number,
    nfceChave: string,
    nfceId: string,
    nfceSituacao: string,
  }
}

export enum CheckoutType { 
  CHECKOUT_TYPE_LINK = 0,
  CHECKOUT_TYPE_ORDER = 1,
}
export enum TipoMovimentacao { 
  TIPO_PEDIDO = 0,
  TIPO_ORCAMENTO = 1,
  TIPO_OS = 2,
  TIPO_AUTORIZACAO = 3,
  TIPO_LICENCA = 4,
}
export enum TipoSituacao { 
  SITUCAO_PEDIDO_PENDENTE = 0,
  SITUCAO_PEDIDO_USADO_EM_MESCLAGEM = 1,
  SITUCAO_PEDIDO_FECHADO = 2,
  SITUCAO_PEDIDO_CANCELADO = 3,
}
export enum TipoMoeda { 
  BRL = 0,
  USD = 1,
}
export enum TabelaPreco { 
  AVISTA = 0,
  APRAZO = 1,
  ATACADO = 2,
}
export enum SituacaoPagamento { 
  SITUCAO_PAGAMENTO_ABERTO = 0,
  SITUCAO_PAGAMENTO_AUTORIZADO = 1,
  SITUCAO_PAGAMENTO_EFETIVADO = 2,
  SITUCAO_PAGAMENTO_LANCADO = 3,
  SITUCAO_PAGAMENTO_CANCELADO = 4,
  SITUCAO_PAGAMENTO_FALHA = 5,
}
export enum Periodicidade { 
  DIA = 0,
  SEMANA = 1,
  MES = 2,
  ANO = 3,
}
export enum Origem { 
  MERCADOLIVRE = 0,
  LOJAVIRTUAL = 1,
  VENDEDOREXTERNO = 2,
  GERAL = 3,
}
export enum TipoItem { 
  PRODUTO = 0,
  SERVICO = 1,
}
export enum TipoDfe { 
  NFE = 0,
  NFCE = 2,
}
