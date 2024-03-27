# Custom Search Attributes

This example shows how custom Search Attributes can be used in your Workflow. This sample uses the Pizza Workflow, and creates a custom Search Attribute called `isOrderFailed` which is a boolean. The user can then use this Search Attribute to query for Workflows where the pizza order has failed.

## Part A: Create a Custom Search Attribute

1. First, you will create your custom Search Attribute, `isOrderFailed`, a boolean type. You can do this in one of your terminals with the following command: `temporal operator search-attribute create --namespace default --name isOrderFailed --type bool`.
2. Make sure you can see all the Search Attributes you have with the command: `temporal operator search-attribute list`. You should now see `isOrderFailed` in this list. It may take up to ten seconds before the attribute appears.

## Part B: Setting a Custom Search Attribute Value While Starting a Workflow

This is not a necessary step. In the `start/main.go` file, you can set Custom Search Attribute by adding them to the options when starting a Workflow execution using [`WorkflowOptions.searchAttributes`](https://docs.temporal.io/dev-guide/go/observability#custom-search-attributes). 

Keep in mind that setting attributes is optional in some attributes' case, and this is more for setting Search Attributes that are known at the start of the Workflow or may not change through the Workflow logic.

## Part C: Upserting Attributes

Within the Pizza Workflow code, we will now dynamically update Search Attributes using [`upsertSearchAttributes`](https://docs.temporal.io/dev-guide/go/observability#upsert-search-attributes).

1. In `workflow.go`, locate `workflow.UpsertSearchAttributes(ctx, failure)` which is used to indicate that the order has not failed. It is marked not failed, because it is in the part of the logic when the Workflow has received the Signal that the order has been fulfilled.

2. In `workflow.go`, locate `workflow.UpsertSearchAttributes(ctx, success)` which is used to indicate that the order has failed. It is marked failed in the part of the Workflow code when the Workflow has received the Signal that the order has not been fulfilled successfully.

## Part D: Running Your Workflow

In one terminal, navigate to the `worker` subdirectory and run `go run main.go`. In another terminal, navigate to the `starter` subdirectory and run `go run main.go`. Now, try your query to retrieve the results.

## Part E: Querying Workflows by Search Attributes

Once you have Workflows tagged with Custom Search Attributes, you can query them based on these attributes. For example, using the CLI:

```shell
temporal workflow list -q 'isOrderFailed=false'
```

This lists all the Workflows that fulfill this query.



### This is the end of the sample.