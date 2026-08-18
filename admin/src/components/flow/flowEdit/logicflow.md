<!-- BEGIN:logicflow-agent-rules -->

# LogicFlow Agent Rules

LogicFlow documentation is available at:

- `node_modules/@logicflow/core/dist/docs/`

Package roles:

- `@logicflow/core`: core graph editor runtime, including canvas, nodes, edges, models, events, rendering, themes, and basic interactions.
- `@logicflow/extension`: official plugins for common product features.
- `@logicflow/layout`: official layout plugins for automatic graph layout.

The docs for `@logicflow/extension` and `@logicflow/layout` are included under:

- `node_modules/@logicflow/core/dist/docs/tutorial/extension/`

Before implementing any LogicFlow feature, check the local docs first to see whether LogicFlow already provides a built-in, extension, or layout capability. If it does, prefer the documented official capability instead of reimplementing it from scratch.

If an official package is needed but not installed, ask the user before installing it.
<!-- END:logicflow-agent-rules -->
