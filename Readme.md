# golang junior dev

##

./cmd/server - ������
./cmd/test_client - ��� �������, ������ ���� ������� ������
./internal/exchanges - ���������� ���������

## ������

### Exchanges

�������� REST ������ �� ������� ���� ��������� ������� ��� ��������� ����� �����. �� ���� ����������� HTTP ������ � �������:

```json
{
  "amount": 400,
  "banknotes": [5000, 2000, 1000, 500, 200, 100, 50]
}
```

���

- **amount** � _����� �����_
- **banknotes** � _��������� �������� �������_

������ ������:

```json
{
  "exchanges": [
    [200, 200],
    [200, 100, 100],
    [200, 100, 50, 50],
    [200, 50, 50, 50, 50],
    [100, 100, 100, 100],
    [100, 100, 100, 50, 50],
    [100, 100, 50, 50, 50, 50],
    [100, 50, 50, 50, 50, 50, 50],
    [50, 50, 50, 50, 50, 50, 50, 50]
  ]
}
```

## ���������� � ����������

- ������������ (����, ����, ������� �����������)
- graceful shutdown
- unit ����� ���������
- �������� � ������������� git �����������
