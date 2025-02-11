O que é Graceful shutdown?
- Garante o processamento de uma requisição até o final, sem interrompe-lo.

os.Signal: sinal do sistema operacional

	signal.Notify(stop, os.Interrupt)
- os.Interrupt: interrupção como ctrl + C
- syscall.SIGTERM
- syscall.SIGINT: interrupção enviada pelo sistema operacional.