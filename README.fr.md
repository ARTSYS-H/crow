<div align="center">
    <img src="./assets/images/logo-crow.png" alt="Crow Logo">
    <h1>Crow</h1>
</div>

[![Go Reference](https://pkg.go.dev/badge/github.com/ARTSYS-H/crow/pkg/crow.svg)](https://pkg.go.dev/github.com/ARTSYS-H/crow/pkg/crow)

Crow est une bibliothèque Go conçue pour créer des applications en ligne de commande de manière simple et intuitive en utilisant les struct fields et les tags. Inspirée par des projets comme [Commandeer](https://github.com/jaffee/commandeer), Crow vise à fournir une solution plus directe et "plug & play" pour la création de petites applications ou scripts, réduisant ainsi la complexité souvent associée à des bibliothèques comme [Cobra](https://github.com/spf13/cobra).

> :warning: **Attention** Crow est toujours en développement. Considérez-la comme expérimental.

## Origine du Projet

Crow a été développé pour répondre à un besoin spécifique : la création rapide et simple de petites applications en ligne de commande. Personnellement, je me suis souvent retrouvé à créer des scripts et petites applications où des bibliothèques comme [Cobra](https://github.com/spf13/cobra) introduisaient une complexité inutile. De plus, des outils comme [Commandeer](https://github.com/jaffee/commandeer), qui partagent une approche similaire, semblent avoir perdu de leur dynamisme en termes de maintenance. Crow est donc né pour offrir une alternative simple et bien maintenue.

## Caractéristiques

- **Simplicité** : Créez des applications CLI avec un minimum de code et de configuration.
- **Utilisation de Structs et Tags** : Définissez vos commandes et options directement dans des structures Go avec des tags.
- **Génération Automatique de Messages d'Aide** : Les messages d'aide pour vos commandes sont générés automatiquement, facilitant la documentation et l'utilisation de votre application.
- **Plug and Play** : Conçu pour être facile à intégrer et à utiliser sans configuration complexe.
- **Idéal pour les Petits Projets** : Parfait pour les scripts et petites applications où des bibliothèques comme [Cobra](https://github.com/spf13/cobra) seraient excessives.

## Installation

Pour installer Crow, utilisez la commande `go get` :

```bash
go get github.com/ARTSYS-H/crow/pkg/crow
```

## Utilisation

Voici un exemple simple pour démarrer avec Crow :
```go
package main

import (
    "github.com/ARTSYS-H/crow/pkg/crow"
)

type MyCommand struct {
    Name string `help:"You Name"`
    Age  int    `help:"Your age"`
}

func (mc *MyCommand) Run() error {
    // Do your stuff here
}

func main() {
    app := crow.New("App Name", "App Description")
    command := &MyCommand{
        Name: "Lucas",
        Age: 27,
    }
    app.AddCommand(command, "Description of the command")
    err = app.Execute(os.Args)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

### Explication

- **Structs et Tags** : Définissez vos commandes et options en tant que champs d'une struct avec des tags pour spécifier les options de ligne de commande.
- **crow.Execute** : Utilisez cette fonction pour analyser les arguments de la ligne de commande.

## Contribution

Les contributions sont les bienvenues ! Si vous avez des suggestions, des corrections de bugs ou des améliorations, n'hésitez pas à ouvrir une issue ou une pull request.

[commandeer]: (https://github.com/jaffee/commandeer)
[cobra]: (https://github.com/spf13/cobra)
