Cobra CLI - para auxiliar a criar projetos com inicialização por linha de comando

Repo: https://github.com/spf13/cobra-cli


go install github.com/spf13/cobra-cli@latest
comando: cobra-cli

Iniciar:
- cobra-cli init
-- vai gerar os arquivos
--- LICENSE
--- main.go
--- cria CMD/root.go

Criar comando:
- cobra-cli add nome_do_comando (é gerado o arquivo cmd/nome_do_comando.go)

Se rodar o main.go agora vai exibir como opção de comando o nome_do_comando

Para rodar o nome comando:
-go run main.go nome_do_comando 

Flags:

Para adicinar uma flag em um comando, alterar o método init() do nome_do_comando.go
exemplo:

    testeCmd.Flags().StringP("comando", "c", "", "descrição da flag")  -- StringP ajuda a criar um flag com o shorthand
    testeCmd.MarkFlagRequire("comando")  -- para tornar a flag obrigatória
    testeCmd.MarkFlagsRequiredTogether("comando", "comando2",...)  -- para tornar a flags obrigatórias

todas essas formas de chamar o comando funcionam:
-> go run main.go nome_do_comando --comando=valor
-> go run main.go nome_do_comando --comando valor
-> go run main.go nome_do_comando -c valor

Comandos encadeados

-> cobra-cli add create -p 'categoryCmd' // o comando create será filho do comando category
-> go run main.go category create


Tipos de Flags

- PersistentFlags
Consegue enxergar o comando na cadeia interia, pois é criado no comando principal.

Na func init() do comando parent
exemplo: categoryCmd.PersistentFlags().String("name", "valor default", "descrição")

comando: to run main.go category            // é exibido o flag name como possibilidade
comando: to run main.go category create     // é exibido o flag name como possibilidade


- Defaults

    testeCmd.Flags().StringP("comando", "c", "", "descrição da flag")  -- StringP ajuda a cria rum flag com o shorthand


Quando se cria uma flag, se define o valor default e se terá shorthand:
- categoryCmd.PersistentFlags().String("name", "valor default", "descrição")
- categoryCmd.PersistentFlags().StringP("name", "n", "valor default", "descrição")

Tipos de flags:
- categoryCmd.PersistentFlags().BoolP("exists", "e", false, "descrição")
- categoryCmd.PersistentFlags().Int32P("id", "i", 0, "descrição")


Para pegar o valor da flag 

    Run: func(cmd *cobra.Command, args []string) {
        name, _ := cmd.Flags().GetString("name")
        exists, _ := cmd.Flags().GetBool("exists")
    }

obs: para flags boleanas, apenas citar a flag já torna seu valor true, ex: go run main teste -e


É possível preencher o valor de uma variável, por referência, com o valor do flag.

var category string
func init() {
    categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "valor default", "descrição")
}


Hooks:
Permitem fazer execuções antes, depois, e outras opções do comando Run() ser rodado, ex:
 Run: func(cmd *cobra.Command, args []string) {
 },
 PreRun: func(cmd *cobra.Command, args []string) {
  },
 PostRun: func(cmd *cobra.Command, args []string) {
  }
 RunE: func(cmd *cobra.Command, args []string) error {
    // Outra forma de executar um comando, diferença é que ele retorna um erro.
 }

Para injetar classes nos comandos:

criar um type no root.go, ex:

type RunEFunc func(cmd *cmd.Command, args []string) error


