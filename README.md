---

<p align="center">
  <strong>
    <a href="https://opentelemetry.io/docs/collector/getting-started/">Getting Started</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/CONTRIBUTING.md">Getting Involved</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://gitter.im/open-telemetry/opentelemetry-service">Getting In Touch</a>
  </strong>
</p>

<p align="center">
  <a href="https://github.com/open-telemetry/opentelemetry-collector-contrib/actions/workflows/build-and-test.yml?query=branch%3Amain">
    <img alt="Build Status" src="https://img.shields.io/github/workflow/status/open-telemetry/opentelemetry-collector-contrib/build-and-test/main?style=for-the-badge">
  </a>
  <a href="https://goreportcard.com/report/github.com/open-telemetry/opentelemetry-collector-contrib">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/open-telemetry/opentelemetry-collector-contrib?style=for-the-badge">
  </a>
  <a href="https://codecov.io/gh/open-telemetry/opentelemetry-collector-contrib/branch/main/">
    <img alt="Codecov Status" src="https://img.shields.io/codecov/c/github/open-telemetry/opentelemetry-collector-contrib?style=for-the-badge">
  </a>
  <a href="https://github.com/open-telemetry/opentelemetry-collector-contrib/releases">
    <img alt="GitHub release (latest by date including pre-releases)" src="https://img.shields.io/github/v/release/open-telemetry/opentelemetry-collector-contrib?include_prereleases&style=for-the-badge">
  </a>
  <img alt="Beta" src="https://img.shields.io/badge/status-beta-informational?style=for-the-badge&logo=data:image/png;base64,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">
</p>

<p align="center">
  <strong>
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/vision.md">Vision</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/design.md">Design</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/monitoring.md">Monitoring</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/performance.md">Performance</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/security.md">Security</a>
    &nbsp;&nbsp;&bull;&nbsp;&nbsp;
    <a href="https://github.com/open-telemetry/opentelemetry-collector/blob/main/docs/roadmap.md">Roadmap</a>
  </strong>
</p>

---

# OpenTelemetry Collector Contrib

This is a repository for OpenTelemetry Collector components that are not suitable for the  [core repository](https://github.com/open-telemetry/opentelemetry-collector) of the collector. 

The official distributions, core and contrib, are available as part of the [opentelemetry-collector-releases](https://github.com/open-telemetry/opentelemetry-collector-releases) repository. Some of the components in this repository are part of the "core" distribution, such as the Jaeger and Prometheus components, but most of the components here are only available as part of the "contrib" distribution. Users of the OpenTelemetry Collector are also encouraged to build their own custom distributions with the [OpenTelemetry Collector Builder](https://github.com/open-telemetry/opentelemetry-collector/tree/main/cmd/builder), using the components they need from the core repository, the contrib repository, and possibly third-party or internal repositories.

Each component has its own support levels, as defined in the following sections. For each signal that a component supports, there's a stability level, setting the right expectations. It is possible then that a component will be **Stable** for traces but **Alpha** for metrics and **In Development** for logs.

## Stability levels

While we intend to provide high-quality components as part of this repository, we acknowledge that not all of them are ready for prime time. As such, each component should list its current stability level for each telemetry signal, according to the following definitions:

### In development

Not all pieces of the component are in place yet and it might not be available as part of any distributions yet. Bugs and performance issues should be reported, but it is likely that the component owners might not give them much attention. Your feedback is still desired, especially when it comes to the user-experience (configuration options, component observability, technical implementation details, ...). Configuration options might break often depending on how things evolve. The component should not be used in production.

### Alpha

The component is ready to be used for limited non-critical workloads and the authors of this component would welcome your feedback. Bugs and performance problems should be reported, but component owners might not work on them right away. The configuration options might change often without backwards compatibility guarantees.

### Beta

Same as Alpha, but the configuration options are deemed stable. While there might be breaking changes between releases, component owners should try to minimize them. A component at this stage is expected to have had exposure to non-critical production workloads already during its **Alpha** phase, making it suitable for broader usage.

### Stable

The component is ready for general availability. Bugs and performance problems should be reported and there's an expectation that the component owners will work on them. Breaking changes, including configuration options and the component's output are not expected to happen without prior notice, unless under special circumstances.

### Deprecated

The component is planned to be removed in a future version and no further support will be provided. Note that new issues will likely not be worked on. When a component enters "deprecated" mode, it is expected to exist for at least two minor releases. See the component's readme file for more details on when a component will cease to exist.

## Gated features

Some features are hidden behind feature gates before they are part of the main code path for the component. Note that the feature gates themselves might be at different [lifecycle stages](https://github.com/open-telemetry/opentelemetry-collector/blob/v0.49.0/service/featuregate/README.md#feature-lifecycle).

## Support

Each component is supported either by the community of OpenTelemetry Collector Contrib maintainers, as defined by the GitHub group [@open-telemetry/collector-contrib-maintainer](https://github.com/orgs/open-telemetry/teams/collector-contrib-maintainer), or by specific vendors. See the individual README files for information about the specific components.

The OpenTelemetry Collector Contrib maintainers may at any time downgrade specific components, including vendor-specific ones, if they are deemed unmaintained or if they pose a risk to the repository and/or binary distribution.

Even though the OpenTelemetry Collector Contrib maintainers are ultimately responsible for the components hosted here, actual support will likely be provided by individual contributors, typically a code owner for the specific component.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

Triagers ([@open-telemetry/collector-contrib-triagers](https://github.com/orgs/open-telemetry/teams/collector-contrib-triagers))
- [Alolita Sharma](https://github.com/alolita), AWS
- [Punya Biswal](https://github.com/punya), Google
- [Steve Flanders](https://github.com/flands), Splunk

Approvers ([@open-telemetry/collector-contrib-approvers](https://github.com/orgs/open-telemetry/teams/collector-contrib-approvers)):

- [Anthony Mirabella](https://github.com/Aneurysm9), AWS
- [David Ashpole](https://github.com/dashpole), Google
- [Dmitrii Anoshin](https://github.com/dmitryax), Splunk
- [Pablo Baeyens](https://github.com/mx-psi), DataDog
- [Przemek Maciolek](https://github.com/pmm-sumo), Sumo Logic

Maintainers ([@open-telemetry/collector-contrib-maintainer](https://github.com/orgs/open-telemetry/teams/collector-contrib-maintainer)):

- [Daniel Jaglowski](https://github.com/djaglowski), observIQ
- [Juraci Paixão Kröhling](https://github.com/jpkrohling), Grafana Labs
- [Alex Boten](https://github.com/codeboten), Lightstep
- [Bogdan Drutu](https://github.com/BogdanDrutu), Splunk

Emeritus Maintainers
- [Tigran Najaryan](https://github.com/tigrannajaryan), Splunk

Learn more about roles in the [community repository](https://github.com/open-telemetry/community/blob/main/community-membership.md).

## PRs and Reviews

When creating a PR please following the process [described
here](https://github.com/open-telemetry/opentelemetry-collector/blob/main/CONTRIBUTING.md#how-to-structure-prs-to-get-expedient-reviews).

News PRs will be automatically associated with the reviewers based on
[CODEOWNERS](.github/CODEOWNERS). PRs will be also automatically assigned to one of the
maintainers or approvers for facilitation.

The facilitator is responsible for helping the PR author and reviewers to make progress
or if progress cannot be made for closing the PR.

If the reviewers do not have approval rights the facilitator is also responsible
for the official approval that is required for the PR to be merged and if the facilitator
is a maintainer they are responsible for merging the PR as well.

The facilitator is not required to perform a thorough review, but they are encouraged to
enforce Collector best practices and consistency across the codebase and component
behavior. The facilitators will typically rely on codeowner's detailed review of the code
when making the final approval decision. 

We recommend maintainers and approvers to keep an eye on the
[project board](https://github.com/orgs/open-telemetry/projects/3). All newly created
PRs are automatically added to this board. (If you don't see the PR on the board you
may need to add it manually by setting the Project field in the PR view).
