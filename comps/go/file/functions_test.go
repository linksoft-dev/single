package file

import "testing"

func TestConvertJsonFieldToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"SingleField", `[{"myField": "value"}]`, `[{"my_field": "value"}]`},
		{"MultipleFields", `[{"myField": "value", "anotherField": "anotherValue"}]`, `[{"my_field": "value", "another_field": "anotherValue"}]`},
		{"NestedFields", `[{"nestedField": {"innerField": "innerValue"}}]`, `[{"nested_field": {"inner_field": "innerValue"}}]`},
		{"AlreadySnakeCase", `[{"my_field": "value"}]`, `[{"my_field": "value"}]`},
		{"SubfieldsWithArrays", `[{"arrayField": [{"subField1": "value1"}, {"subField2": "value2"}]}]`, `[{"array_field": [{"sub_field1": "value1"}, {"sub_field2": "value2"}]}]`},
		{"SameFieldName", `[{"fieldName": "value", "nestedField": {"fieldName": "value"}}]`, `[{"field_name": "value", "nested_field": {"field_name": "value"}}]`},
		{"SameFieldName", `[{"id": "2fbVPK8UhYOVgEWBpU2MltiOVZY", "ncm": "08031000", "nome": "TESTE", "userId": "1zjCIKl7nYAkEJKzPYiXXYqN68k", "actions": null, "userName": "RODRIGO", "createdAt": "2024-04-25T15:28:58.163560694-03:00", "updatedAt": "2024-04-25T15:29:40.777269703-03:00", "tributacaoId": "1zlkyMOIcKkI5H53adJ7emZ59jM", "tributacao_nome": "SUBSTITUICAO TRIBUTARIA"}]`, `[{"id": "2fbVPK8UhYOVgEWBpU2MltiOVZY", "ncm": "08031000", "nome": "TESTE", "user_id": "1zjCIKl7nYAkEJKzPYiXXYqN68k", "actions": null, "user_name": "RODRIGO", "created_at": "2024-04-25T15:28:58.163560694-03:00", "updated_at": "2024-04-25T15:29:40.777269703-03:00", "tributacao_id": "1zlkyMOIcKkI5H53adJ7emZ59jM", "tributacao_nome": "SUBSTITUICAO TRIBUTARIA"}]`},
		{"ValueInCamelCase", `[{"myField": "myValueInCamelCase"}]`, `[{"my_field": "myValueInCamelCase"}]`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertJsonFieldToSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("ConvertJsonFieldToSnakeCase(%s) = %s; expected %s", tt.input, result, tt.expected)
			}
		})
	}
}
