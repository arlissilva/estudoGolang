package contextEstudos

import (
	"context"
	"time"
)

func Entrypoint() {
	//context.Background() cria um context vazio, sem regras ou valores associados.
	// Ele é usado como ponto de partida para criar contextos derivados.
	ctx := context.Background()

	//Aqui criamos um novo context com timeout de 3 segundos usando context.WithTimeout.
	// O defer cancel() garante que a função cancel seja chamada para liberar os recursos
	// associados ao context quando a função Entrypoint terminar, por boa pratica sempre devemos cancelar
	// o context para evitar vazamento de recursos, mesmo que o timeout seja atingido ou a operação seja concluída.
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	bookingHotel(ctx)
}

func bookingHotel(ctx context.Context) {
	//Aqui usamos um select para aguardar a conclusão da reserva do hotel ou o cancelamento do context.
	// Se o context for cancelado (por exemplo, se o timeout for atingido), o case <-ctx.Done() será executado,
	// imprimindo a mensagem de erro e cancelando a reserva do hotel. Se a reserva do hotel for concluída antes do timeout,
	// o case <-time.After(5 * time.Second) será executado, indicando que a reserva foi concluída com sucesso.
	select {
	case <-ctx.Done():
		println(ctx.Err())
		println("Cancelando a reserva do hotel...")
		return
	case <-time.After(5 * time.Second):
		println("Reserva do hotel concluída!")
	}
}
