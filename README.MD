Clean Architecture

CQRS

Singleton

Repository

Factory

Strategy

Middleware

Dependency Injection

Unit of Work

Domain Event

Adapter pattern (repo impl)

--
┌────────────────────────┐
│ Delivery │
│ (HTTP / Router) │
└────────────▲───────────┘
│
┌────────────┴───────────┐
│ Application │
│ (CQRS - UseCases) │
└────────────▲───────────┘
│
┌────────────┴───────────┐
│ Domain │
│ (Entities + Rules) │
└────────────▲───────────┘
│
┌────────────┴───────────┐
│ Infrastructure │
│ (DB, Logger, Redis) │
└────────────────────────┘

--
/cmd
/api
main.go

/internal
/config
config.go

    /domain
        /entity
            user.go
            article.go
        /repository
            user_repository.go
            article_repository.go
        /service
            password_hasher.go

    /application
        /command
            /user
                create_user.go
            /article
                create_article.go
                update_article.go
                delete_article.go
        /query
            /user
                get_user.go
            /article
                list_articles.go
                get_article.go

    /infrastructure
        /database
            db.go
            unit_of_work.go
        /repository
            user_repository_impl.go
            article_repository_impl.go
        /logger
            logger.go
        /auth
            jwt.go

    /delivery
        /http
            router.go
            /handler
                user_handler.go
                article_handler.go
            /middleware
                auth_middleware.go
                logging_middleware.go

/pkg
/singleton
container.go

/go.mod
