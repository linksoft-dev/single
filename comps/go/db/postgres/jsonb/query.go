package jsonb

import (
	"golang.org/x/exp/maps"
	"reflect"
	"strings"
)

type join struct {
	sourceTable string
	targetTable string
	sourceField string
	targetField string
}

//
//// Or Funcao para aplicar condicional OR ao SQL
//func (q *db.Query) Or(field string, conditions ...string) *db.Query {
//	if !strings.Contains(field, "doc") {
//		field = fmt.Sprintf("doc->>'%s'", field)
//	}
//	where := ""
//	for _, c := range conditions {
//		if c == "" {
//			continue
//		}
//		condition := c
//		if !strings.Contains(condition, "=") && !strings.Contains(condition, "is") {
//			condition = fmt.Sprintf(" = '%s' ", condition)
//		}
//		if where == "" {
//			where = "("
//		}
//		where = fmt.Sprintf("%s %s %s OR ", where, field, condition)
//	}
//	where = strings.TrimSuffix(where, "OR ")
//	if where != "" {
//		where += ")"
//		q.AddWhere(where)
//	}
//	return q
//}
//
//func (q *db.Query) Join(sourceTable, sourceField, targetTable, targetField string) *db.Query {
//	q.join = append(q.join, join{sourceTable: sourceTable, sourceField: sourceField,
//		targetTable: targetTable, targetField: targetField})
//	return q
//}
//
//func (q *db.Query) JoinString(v string) *db.Query {
//	q.joinString += v
//	return q
//}
//
//func (q *db.Query) Raw(v string, params ...interface{}) *db.Query {
//	q.rawQuery += v
//	q.rawParams = params
//	return q
//}
//
//func (q *db.Query) Find2(dest interface{}) (err error) {
//
//	// se estiver consultado com rawquery, nao processe nada, apenas faça o scan para o `dest`
//	if q.rawQuery == "" {
//
//		// sql fixa para todas a querys, dando prioridade a filtragem com softdelete e tabela/collection
//		sql := "deleted_at is null and "
//		if strings.Contains(q.TableName, "%") {
//			sql = fmt.Sprintf("%s collection like '%s%%'", sql, q.TableName)
//		} else {
//			sql = fmt.Sprintf("%s collection='%s'", sql, q.TableName)
//		}
//		q.fixedWhere = sql
//
//		// caso nao tenha passado o campo doc, adicione automaticamente
//		for idx, value := range q.conditions {
//			if !strings.Contains(value.Field, "doc") {
//				q.conditions[idx].Field = fmt.Sprintf("doc ->> '%s'", value.Field)
//			}
//		}
//		if q.TableName == "" && q.rawQuery == "" {
//			return errors.New("Nome da tabela nao foi passado para a Query")
//		}
//
//		q.From("org_" + q.db.TenantId)
//
//		// se os fields solicitados for diferente do padrão que é *,
//		//então o resultado é customizado e nao precisa converter para a estrutura de array de Doc
//		if len(q.fields) > 0 && q.fields[0] != "*" {
//			err = q.FindRaw(dest)
//			if err != nil {
//				if q.db.createTableIfDoesntExists(err) {
//					return q.Find2(dest)
//				}
//			}
//			return
//		}
//
//		//Se o Sort n ta com o cast pro doc, o mesmo deve ser adicionado
//		for i, v := range q.Sort {
//			if !strings.Contains(i, "doc") {
//				delete(q.Sort, i)
//				q.Sort[fmt.Sprintf("doc -> '%s'", i)] = v
//			}
//		}
//		// caso seja uma pesquisa padrao, ou seja, passou da condicao acima, então adicione a ordenacao pelo ultimo inserido
//		// caso nao tenha nenhuma instrucao de sort
//		if len(q.Sort) == 0 {
//			q.OrderByDesc("doc -> 'createdAt'")
//		}
//	}
//
//	// concatene os docs em uma string para formar um array de dos em string, para então fazer o marshal para o destino
//	docs := []Doc{}
//	if err = q.FindRaw(&docs); err != nil {
//		if q.db.createTableIfDoesntExists(err) {
//			return q.Find2(dest)
//		}
//		return
//	}
//
//	// se está trazendo o primeiro registro ou o ultimo, entao ja faça o parse direto com a variavel `dest`
//	if q.first == 1 || q.last == 1 {
//		if len(docs) > 0 {
//			err = json.Unmarshal([]byte(docs[0].Doc), dest)
//		}
//		return
//	}
//
//	var sb strings.Builder
//	sb.WriteString("[")
//	for _, value := range docs {
//		sb.WriteString(value.Doc)
//		sb.WriteString(",")
//	}
//	str := sb.String()
//	str = strings.TrimSuffix(str, ",")
//	str += "]"
//	err = json.Unmarshal([]byte(str), dest)
//
//	return
//}

// getChildTagsByStruct retorna as tags de uma struct necessarias a consulta de um filho de um registro
func getChildTagsByStruct(tp reflect.Type) (fatherFields map[string]string, childFields map[string]string) {
	fatherFields = map[string]string{}
	childFields = map[string]string{}
	for i := 0; i < tp.NumField(); i++ {
		f := tp.Field(i)
		if f.Tag.Get("fatherField") != "" {
			fatherFields[f.Tag.Get("fatherField")] = strings.Replace(f.Tag.Get("json"), ",omitempty", "", 1)
			continue
		}
		if f.Tag.Get("json") == "" {
			if tp.Field(i).Type.Kind() == reflect.Struct {
				newFatherFields, newChildFilds := getChildTagsByStruct(tp.Field(i).Type)
				//adiciona os dados obtidos na recursividade
				maps.Copy(fatherFields, newFatherFields)
				maps.Copy(childFields, newChildFilds)
				continue
			}
			continue
		}
		//se for encontrado traço nessas tags o campo deve ser ignorado
		if f.Tag.Get("json") == "-" || f.Tag.Get("childType") == "-" {
			continue
		}
		//por padrao o tipo é text
		typeChild := "text"
		if f.Tag.Get("childType") != "" {
			typeChild = f.Tag.Get("childType")
		}
		childFields[strings.Replace(f.Tag.Get("json"), ",omitempty", "", 1)] = typeChild
	}
	return
}

// // FindChild realiza a consulta de uma registro filho de um registro principal e retorna a lista dos mesmos
//
//	func (q *db.Query) FindChild(childName string, newStructure, dest interface{}) (err error) {
//		// se estiver consultado com rawquery, nao processe nada, apenas faça o scan para o `dest`
//		if q.rawQuery == "" {
//			if childName == "" {
//				return errors.New("É necessário informar a TAG JSON do filho usado na structure Pai.")
//			}
//
//			// sql fixa para todas a querys, dando prioridade a filtragem com softdelete e tabela/collection
//			sql := "deleted_at is null and "
//			if strings.Contains(q.TableName, "%") {
//				sql = fmt.Sprintf("%s collection like '%s%%'", sql, q.TableName)
//			} else {
//				sql = fmt.Sprintf("%s collection='%s'", sql, q.TableName)
//			}
//			q.fixedWhere = sql
//
//			// caso nao tenha passado com o indicador child ou father, adicione o parse child como padrao
//			// obs o indicador pode sert father, neste caso tem q ser informado manualmente
//			for idx, value := range q.conditions {
//				if !strings.Contains(value.Field, "child") && !strings.Contains(value.Field, "father") &&
//					!strings.Contains(value.Field, "doc") {
//
//					q.conditions[idx].Field = fmt.Sprintf("child.%s", value.Field)
//				}
//			}
//			if q.TableName == "" && q.rawQuery == "" {
//				return errors.New("Nome da tabela nao foi passado para a Query")
//			}
//
//			q.From("org_" + q.db.TenantId)
//
//			// se os fields solicitados for diferente do padrão que é *,
//			//então o resultado é customizado e nao precisa converter para a estrutura de array de Doc
//			if len(q.fields) > 0 && q.fields[0] != "*" {
//				err = q.FindRaw(dest)
//				if err != nil {
//					if q.db.createTableIfDoesntExists(err) {
//						return q.Find2(dest)
//					}
//				}
//				return
//			}
//
//			//obtem as tags necessárias para formar o SQL de consulta
//			itemReflectVOf := reflect.ValueOf(newStructure)
//			itemType := reflect.Indirect(itemReflectVOf).Type()
//			if itemType.Kind() != reflect.Struct {
//				return fmt.Errorf("newStructure is not a struct.")
//			}
//			fatherFields, childFields := getChildTagsByStruct(itemType)
//
//			//Campos que ficarão dentro do SELECT para serem retornados
//			selectStr := ""
//			//Campos do filho com parse do tipo para ser usado no FROM
//			childFrom := ""
//			for x, k := range fatherFields {
//				selectStr = fmt.Sprintf(`%s father.doc -> %s as "%s" ,`, selectStr, x, k)
//			}
//			if len(childFields) == 0 {
//				selectStr = strings.TrimSuffix(selectStr, ",")
//			}
//			indice := 0
//			for c, v := range childFields {
//				selectStr = fmt.Sprintf(`%s child."%s" as "%s" ,`, selectStr, c, c)
//				childFrom = fmt.Sprintf(`%s "%s" %s,`, childFrom, c, v)
//				if indice == len(childFields)-1 {
//					selectStr = strings.TrimSuffix(selectStr, ",")
//					childFrom = strings.TrimSuffix(childFrom, ",")
//				}
//				indice++
//			}
//
//			// caso nao tenha nenhuma instrucao de sort ordene decrescentemente pela data de criacao do pai
//			if len(q.Sort) == 0 {
//				q.OrderByDesc("father.doc -> 'createdAt'")
//			}
//
//			//obtem os filtros da consulta
//			whereSql, params := q.getWhere(1)
//			if whereSql != "" {
//				whereSql = fmt.Sprintf("WHERE  (doc ->> '%s' is not null) AND %s",
//					childName, whereSql)
//			}
//
//			// ordenacao
//			orderBy := ""
//			limit := ""
//			offset := ""
//			if q.last > 0 {
//				orderBy = `order by child."createdAt" desc limit 1`
//			} else if q.first > 0 {
//				orderBy = `order by child."createdAt" asc limit 1`
//			} else {
//				for key, value := range q.Sort {
//					orderBy += fmt.Sprintf("%s %s,", key, value)
//				}
//				if orderBy != "" {
//					orderBy = fmt.Sprintf("order by %s ", strings.TrimSuffix(orderBy, ","))
//				}
//
//				if q.limit > 0 {
//					limit = fmt.Sprintf("limit %d", q.limit)
//				}
//
//				if q.page > 1 {
//					offset = fmt.Sprintf("offset %d", (q.page-1)*q.limit)
//				}
//			}
//
//			//forma a instrucao SQL com os dados obtidos
//			sql = fmt.Sprintf(`
//			SELECT row_to_json(doc)::jsonb as doc FROM (
//				SELECT
//					%s
//				FROM %s as father, jsonb_to_recordset(father.doc -> '%s') as child(
//					%s
//				)
//				 %s
//				%s %s %s
//			) as doc`,
//				selectStr, q.TableName, childName, childFrom, whereSql, orderBy, limit, offset)
//			//realiza a consulta pelas funcoes padão
//			err = q.Raw(sql, params...).Find2(dest)
//			return
//		}
//
//		q.Raw(q.rawQuery).Find2(dest)
//		return
//	}
//
// // FindChildFirst traz o primeiro documento da condicao, sempre traz somente um registro
//
//	func (q *db.Query) FindChildFirst(childName string, newStructure, dest interface{}) (err error) {
//		q.First(1)
//		err = q.FindChild(childName, newStructure, dest)
//		return
//	}
//
//	func (q *db.Query) Find() error {
//		q.Eq("deleted_at", nil)
//		q.Eq("collection", q.TableName)
//		q.From("org_" + q.db.TenantId)
//		return q.FindRaw(nil)
//	}
//
//	func (q *db.Query) FindWithoutOrg() error {
//		q.Eq("deleted_at", nil)
//		return q.FindRaw(nil)
//	}
//
// // FindRaw essa funcao pega os registros ignorando se o mesmo foi deletado ou nao,
// // usada também para trazer selects personalizados
//
//	func (q *db.Query) FindRaw(dest interface{}) (err error) {
//		sqlQuery, params, err := q.GetSqlQueryWithParams()
//		if err != nil {
//			return err
//		}
//		return q.db.Select(dest, sqlQuery, params...)
//	}
//
// // FindRaw essa funcao pega os registros ignorando se o mesmo foi deletado ou nao
//
//	func (q *db.Query) FindOneRaw(dest interface{}) (err error) {
//		sqlQuery, params, err := q.GetSqlQueryWithParams()
//		if err != nil {
//			return err
//		}
//		return q.db.Select(dest, sqlQuery, params...)
//	}
//
//func (q *db.Query) SetAppQueryCriteria(filter db.Query) {
//
//}
