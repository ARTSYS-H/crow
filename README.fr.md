<div align="center">
    <img src="./assets/images/logo-crow.png" alt="Crow Logo">
    <h1>Crow</h1>
</div>

Crow est une bibliothèque Go conçue pour créer des applications en ligne de commande de manière simple et intuitive en utilisant les struct fields et les tags. Inspirée par des projets comme [Commandeer](commandeer), Crow vise à fournir une solution plus directe et "plug & play" pour la création de petites applications ou scripts, réduisant ainsi la complexité souvent associée à des bibliothèques comme [Cobra](cobra).

## Origine du Projet

Crow a été développé pour répondre à un besoin spécifique : la création rapide et simple de petites applications en ligne de commande. Personnellement, je me suis souvent retrouvé à créer des scripts et petites applications où des bibliothèques comme [Cobra](cobra) introduisaient une complexité inutile. De plus, des outils comme [Commandeer](commandeer), qui partagent une approche similaire, semblent avoir perdu de leur dynamisme en termes de maintenance. Crow est donc né pour offrir une alternative simple et bien maintenue.

## Caractéristiques

- **Simplicité** : Créez des applications CLI avec un minimum de code et de configuration.
- **Utilisation de Structs et Tags** : Définissez vos commandes et options directement dans des structures Go avec des tags.
- **Génération Automatique de Messages d'Aide** : Les messages d'aide pour vos commandes sont générés automatiquement, facilitant la documentation et l'utilisation de votre application.
- **Plug and Play** : Conçu pour être facile à intégrer et à utiliser sans configuration complexe.
- **Idéal pour les Petits Projets** : Parfait pour les scripts et petites applications où des bibliothèques comme [Cobra](cobra) seraient excessives.

## Installation

Pour installer Crow, utilisez la commande `go get` :

```bash
go get github.com/ARTSYS-H/crow
```

## Utilisation

Voici un exemple simple pour démarrer avec Crow :
```go
package main

import (
    "github.com/ARTSYS-H/crow"
)

type MyCommand struct {
    Name string `crow:"name,n,required,help:Your name"`
    Age  int    `crow:"age,a,help:Your age"`
}

func main() {
    var cmd MyCommand
    crow.Parse(&cmd)

    // Utilisez cmd.Name et cmd.Age ici
    println("Name:", cmd.Name)
    println("Age:", cmd.Age)
}
```

### Explication

- **Structs et Tags** : Définissez vos commandes et options en tant que champs d'une struct avec des tags pour spécifier les options de ligne de commande.
- **crow.Parse** : Utilisez cette fonction pour analyser les arguments de la ligne de commande et remplir votre struct.

## Contribution

Les contributions sont les bienvenues ! Si vous avez des suggestions, des corrections de bugs ou des améliorations, n'hésitez pas à ouvrir une issue ou une pull request.

[commandeer]: (https://github.com/jaffee/commandeer)
[cobra]: (https://github.com/spf13/cobra)
