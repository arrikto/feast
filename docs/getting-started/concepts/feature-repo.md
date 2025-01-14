# Feature Repository

## Feature Repo

A feature repository is the collection of python files that define entities, feature views and data sources. Feature Repos also have a `feature_store.yaml` file at their root. 

Users can collaborate by making and reviewing changes to Feast object definitions (feature views, entities, etc) in the feature repo.
But, these objects must be applied, either through API, or the CLI, for them to be available by downstream Feast actions (such as materialization, or retrieving online features). Internally, Feast only looks at the registry when performing these actions, and not at the feature repo directly.

## Declarative Feature Definitions

When using the CLI to apply changes (via `feast apply`), the CLI determines the state of the feature repo from the source files and updates the registry state to reflect the definitions in the feature repo files.
This means that new feature views are added to the registry, existing feature views are updated as necessary, and Feast objects removed from the source files are deleted from the registry. 