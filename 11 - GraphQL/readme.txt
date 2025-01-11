gqlgen.com 
- ferramenta que ajuda a criar o servidor 

depois de gerar os modelos corretos no arquivo schema.graphqls, rodar o comando abaixo para gerar os modelos corretos na pasta model.
go run github.com/99designs/gqlgen generate

A jogada é usar o resolver (graph/resolve.go), que é inicialmente vazio, ter acesso aos DBs dos modelos, para que as mutationResolve tenham acesso aos dados. 

Jogamos as models de Cateogry Course em aquivos separados na pasta model
alteramos o gqgen.yml, na seção models, incluindo o seguinte:

models:
  Category:
    model:
      - renebizelli/graphql/graph/model.categoryategory
  Course:
    model:
      - renebizelli/graphql/graph/model.Course
  ...

rodar o "go run github.com/99designs/gqlgen generate" novamente

Com isso, no shemaresolvers.go foi adicionada as linha:
- func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
- type categoryResolver struct{ *Resolver }

implementar o nome método que foi criado 
	r.CourseDB.FindByCategoryId(*&obj.ID) e usa-lo método Courses(ctx context.Context, obj *model.Category) 


- Para rodar no playground

mutation categoryMutation {
  createCategory(input:{name:"Tecnologia",description: "Cussos de tecnologia"}) {
    id
    name
    description
  }
}

mutation courseMutation {
  createCourse(input:{name:"Go",description: "Curso de GoLang", categoryId: "da1f362f-a319-4380-923f-fecb11c88f41"}) {
    id
    name
    description
    category{id}
  }
}

query categoryQueries {
  categories {
    id
    name
  }
}


query categoryQueriesWithCourses {
  categories {
    id
    name
    courses {
      id
      name
    }
  }
}

query courseQueries {
  courses {
    id
    name
    category{id}
  }
}

VS Extension
GraphQL: Language Feature Support - GraphQL Foundation - graphql.org
