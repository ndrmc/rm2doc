# Rm2Doc
Relational to Document - A tool to convert CATS transactional records in PostgreSQL database to a MongoDB document for use in reports and dashboards

# Building and running
## Dependencies
Rm2Doc depends on the following external libraries in addtion to the standard go libs:
- [GORM](https://github.com/jinzhu/gorm)
- [pq](https://github.com/lib/pq)

## Managing dependencies
Rm2Doc uses [dep](https://github.com/golang/dep) package manager and vendoring tool.
- Run dep ensure to ensure vendor/ is in the correct state and all dependencies are in place

## Building
Rm2Doc builds two targets - a command line tool and a REST API server. More detail coming soon...




