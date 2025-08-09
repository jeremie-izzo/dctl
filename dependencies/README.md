# Dependencies

This directory contains infrastructure dependencies organized by type and version for better scalability and future Kubernetes migration.

## Structure

```
dependencies/
├── datastore/
│   └── v1/
│       ├── compose/
│       │   └── docker-compose.yaml
│       └── k8s/
│           └── README.md
├── elasticsearch/
│   └── v8.17.0/
│       ├── compose/
│       │   └── docker-compose.yaml
│       └── k8s/
│           └── README.md
├── mysql/
│   └── v8.0/
│       ├── compose/
│       │   └── docker-compose.yaml
│       └── k8s/
│           └── README.md
├── pubsub/
│   └── v1/
│       ├── compose/
│       │   └── docker-compose.yaml
│       └── k8s/
│           └── README.md
├── redis/
│   └── v7/
│       ├── compose/
│       │   └── docker-compose.yaml
│       └── k8s/
│           └── README.md
└── temporal/
    └── v1.27.2/
        ├── compose/
        │   └── docker-compose.yaml
        └── k8s/
            └── README.md
```

## Versioning Strategy

Each dependency type has its own versioning scheme:

- **datastore**: Uses semantic versioning (v1, v2, etc.)
- **elasticsearch**: Uses Elasticsearch version numbers (8.17.0, 8.18.0, etc.)
- **mysql**: Uses MySQL version numbers (8.0, 8.1, etc.)
- **pubsub**: Uses semantic versioning (v1, v2, etc.)
- **redis**: Uses Redis version numbers (7, 8, etc.)
- **temporal**: Uses Temporal version numbers (1.27.2, 1.28.0, etc.)

## Adding New Versions

To add a new version of a dependency:

1. Create a new version directory under the appropriate type
2. Copy the existing `compose/docker-compose.yaml` and modify as needed
3. Create a `k8s/` directory with migration notes
4. Update any version-specific configurations
5. Update the main application to reference the new version

## Future Kubernetes Migration

The structure is designed to support easy migration to Kubernetes:

- Each dependency type can have its own Kubernetes manifests
- Version-specific configurations can be maintained separately
- The same versioning strategy applies to both Docker Compose and Kubernetes
- Common configurations can be shared between deployment types

## Usage

To use a specific version of a dependency, reference the path:
```
dependencies/{type}/{version}/compose/docker-compose.yaml
```

Example:
```
dependencies/mysql/v8.0/compose/docker-compose.yaml
```

## Path Structure

The new structure follows this pattern:
- `dependencies/{dependency-type}/{version}/compose/docker-compose.yaml` - Docker Compose configuration
- `dependencies/{dependency-type}/{version}/k8s/` - Kubernetes manifests (future) 