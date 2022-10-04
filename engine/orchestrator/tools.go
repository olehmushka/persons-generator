package orchestrator

import (
	"persons_generator/core/storage"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func parseFindOpts(opts storage.PaginationSortingOpts) *options.FindOptions {
	findOpt := options.Find()

	// sorting
	sort := bson.D{}
	for _, sortOpt := range opts.Sorting {
		sort = append(sort, bson.E{Key: sortOpt.FieldName, Value: getSortDESC(sortOpt.IsDESC)})
	}
	findOpt.SetSort(sort)
	// pagination
	findOpt.SetSkip(int64(opts.Pagination.Offset))
	findOpt.SetLimit(int64(opts.Pagination.Limit))

	return findOpt
}

func getSortDESC(isDESC bool) int {
	if isDESC {
		return -1
	}

	return 1
}
