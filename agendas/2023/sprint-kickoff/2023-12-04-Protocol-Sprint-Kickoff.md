# Team Wins 🎉

- Merged last Cadence 1.0 feature
- Completed merging of all major Storehouse M1 PRs to flow-go master (behind a feature flag) & starting to benchmark storehouse on Testnet.
- Deployed [multiple security fixes](https://github.com/onflow/cadence/pull/2955) on Mainnet


### Mainnet Uptime SLO - Last 14 days (11/10 to 11/24)

|                         | Target | Current Score | Error budget used |
|:------------------------|:------:|:-------------:|:-----------------:|
| Collection Finalization | 99.9%  |    100%       |       0%          |
| Block Finalization      | 99.9%  |    100%       |       0%          |
| Transaction Execution   | 99.9%  |    100%       |       0%          |
| Block Sealing           | 99.9%  |    100%       |       0%          |
| Access API Liveness     | 99.9%  |    100%       |       0%          |

#### Incidents
- Degraded Performance - 11/23 9:30 AM Pacific to 3:30 PM Pacific
  - Collection finalization rate dropped and was choppy.


### **Performance Pod Sprint Objective - Jan B**

**Done last sprint**

Storehouse M1
- [Fix getting highest executed block ID when storehouse is enabled](https://github.com/onflow/flow-go/pull/5109)
- [Fix storage snapshot](https://github.com/onflow/flow-go/pull/5107)
- [Optimize finalized reader to cache last finalized height](https://github.com/onflow/flow-go/pull/5056)
- [Adding flag to enable storehouse](https://github.com/onflow/flow-go/issues/5054)
- [Optimization: Reduce finalized block call](https://github.com/onflow/flow-go/pull/5053)

EVM
- [Events emitted from EVM are not encoded as other FVM events](https://github.com/onflow/flow-go/issues/5079)
  - [Emit events as Cadence events](https://github.com/onflow/flow-go/pull/5090)
- [Add a transaction test](https://github.com/onflow/flow-emulator/issues/515)
- [Transaction execution fails](https://github.com/onflow/flow-go/issues/5068)
  - [Fix the setup process](https://github.com/onflow/flow-go/pull/5069)
- [Add EVM transactions to FVM benchmark tests](https://github.com/onflow/flow-go/pull/5061)

Atree register inlining
- Run migration program modified to use atree inlining (inlined checkpoint files reduced migration peak RAM by ~330GB and inlined checkpoint file is 215GB vs 348GB)
- [Add feature to support mutation of primitive values returned by iterators](https://github.com/onflow/atree/issues/356)
- [Handle edge cases needed to support mutation of inlined maps and arrays during iteration](https://github.com/onflow/atree/issues/356)
  - Ready for review: [Add feature to support mutation for array and map iterators](https://github.com/onflow/atree/pull/359)


Other Improvements
- E-E tests
  - [Fix storage fees test](https://github.com/onflow/flow-e2e-tests/issues/50)
  - [Update e2e tests with new flow-go](https://github.com/onflow/flow-e2e-tests/issues/49)
- Emulator
  - [Unify core contracts address definitions](https://github.com/onflow/flow-go/pull/5033)
  - [Use core contract list from flow-go](https://github.com/onflow/flow-emulator/issues/513)


**This sprint**

- Atree register inlining
  - Merge the [Atree inlining integration with Cadence](https://github.com/onflow/cadence/pull/2882)
  - Continue work on validating [migration](https://github.com/onflow/flow-go/pull/4633) of [integrated solution for Atree register inlining](https://github.com/onflow/cadence/issues/2809)
  - Refactor the migration using [mutable iterator](https://github.com/onflow/atree/pull/359)
- Continue testing / benchmarking of [Storehouse first milestone](https://github.com/onflow/flow-go/issues/4682) (execution state on disk) on Testnet.
- EVM support & development

**On Hold**

- Plan removal of concurrent storage bottlenecks
    - Transaction fee deduction
    - [Cadence Type checker is not reentrant](https://dapperlabs.slack.com/archives/CG0B7CJAJ/p1684434997197079) (type comparison depends on consistent pointer used by programs cache, program cache needs to always return the same pointer to the same type)
- [Execution stack refactor - clear separation of ingestion engine and block computer](https://github.com/onflow/flow-go/issues/4077)
- [Automated Performance Tests](https://github.com/onflow/flow-go/issues/3548)

**Active Epics**

- [Atree register inlining](https://github.com/onflow/atree/issues/292)

### Cadence

### **Stable Cadence - Jan B**
Objective: long-term support release of Cadence with no expected breaking changes

**Done last sprint**

Feature
- [New Behavior for Attachments and Entitlements](https://github.com/onflow/cadence/issues/2921)

FLIP
- [FLIP for new behavior for attachments with entitlements](https://github.com/onflow/flips/issues/213)

Improvements
- 1.0: [References to references are not properly checked](https://github.com/onflow/cadence/issues/2873)
  - [throw error on the creation of a nested reference](https://github.com/onflow/cadence/pull/2965)
- 1.0: [Generalize the migrations and make common codes re-usable](https://github.com/onflow/cadence/issues/2942)
- master: [Improve resource-reference tracking](https://github.com/onflow/cadence/pull/2958)

1.0 Migrations
- [Add String normalizing migration](https://github.com/onflow/cadence/issues/2937)
- [Add account type migration](https://github.com/onflow/cadence/issues/2923)

Bugfixes
- [Port bug fixes from v0.42.5-patch.1](https://github.com/onflow/cadence/issues/2956)

Tech-debt
- [Remove unsafeRandom](https://github.com/onflow/cadence/issues/2909)

Tests
- [Add test for container mutation while iterating](https://github.com/onflow/cadence/issues/2960)

Docs
- [Remove references to destructors and add docs for default destroy events](https://github.com/onflow/cadence-lang.org/issues/31)
- [Rewrite entitlements and attachments docs for new changes](https://github.com/onflow/cadence-lang.org/issues/30)
- [Document Account.Capabilities.exists](https://github.com/onflow/cadence-lang.org/issues/29)
- [Remove incorrect statements about reentrancy](https://github.com/onflow/cadence-lang.org/issues/28)

Updating Downstream dependencies
- SDK [Update cadence to v0.42.6](https://github.com/onflow/flow-go-sdk/issues/529)
- flow-go [Update to Cadence v0.42.5-patch.1](https://github.com/dapperlabs/flow-go/issues/6914)

**This sprint**
- Merge Cadence 1.0 feature branch to Cadence master
- Release Cadence 1.0 RC1
- Release Emulator based on Cadence 1.0 RC1 release
- Continue support EVM on FLow initiative.
- Continue work on Cadence 1.0 migrations.
- Continue Stable Cadence Docs update and knocking tasks off the [tech debt list](https://github.com/onflow/cadence/issues/2642)
- Publish Cadence 1.0 release plan on forum
 
**On Hold**
- Discussion of the re-entrancy edge cases

**Active Epics**
- [Stable Cadence (aka Cadence 1.0)](https://github.com/onflow/cadence/issues/2642)


### **Access & Data Availability - Peter A**
Objective: Make execution data and script execution available on Edge nodes.

**Done last sprint**


**This sprint**

- Script Execution on ANs
  - Use version beacon to ensure correct version for script exec - [Issue 5040](https://github.com/onflow/flow-go/issues/5040)
  - Add GetRegisters API endpoint to ExecutionData API - [Issue 4756](https://github.com/onflow/flow-go/issues/4756)
  - Work with QuickNode to setup script exec in compare mode on public ANs
- Misc
  - Work with 4d on getting event streaming and CCF into libraries.

**Active Epics**

- Script Execution on Access Node - [Issue 4637](https://github.com/onflow/flow-go/issues/4637)
- Integrate local execution state indexes into Access API - [Issue 4750](https://github.com/onflow/flow-go/issues/4750)


### **Permissionless Network - Yahya H**

**Done last sprint**


**Still in Progress**
  - [mainnet24` peer scoring incident](https://github.com/dapperlabs/flow-go/issues/6913)
  - Part-3: [Caching GossipSub Application Specific Score](https://github.com/onflow/flow-go/pull/5045) of [Optimizing memory-intensive RPC inspection operations](https://github.com/dapperlabs/flow-go/issues/6870)
  - [Part-2 and -3 of the long term fix for AN-LN peer blocking issue on mainnet23](https://github.com/dapperlabs/flow-go/issues/6895)
  - [GossipSub Message Replay Attack](https://github.com/dapperlabs/flow-go/issues/6809) [PR5058](https://github.com/onflow/flow-go/pull/5058)

**Starting Next Sprint**:
  - [Concluding GossipSub Message Forensics FLIP and planning the development](https://github.com/onflow/flips/pull/195)
  - [[BFT Testing] Refactor Orchestrator lock contension to use worker pools](https://github.com/dapperlabs/flow-go/issues/6884)
  - [Determining an appropriate retention rate for historical scoring data](https://github.com/dapperlabs/flow-go/issues/6466)
  - [Decision Making for Persisting or Non-Persisting Spamming Records of Peers in GossipSub](https://github.com/dapperlabs/flow-go/issues/6663)


**Active Epics**

- https://github.com/dapperlabs/flow-go/issues/6287
- https://github.com/dapperlabs/flow-go/issues/6468
- BFT https://github.com/dapperlabs/flow-go/issues/6142
- BFT https://github.com/dapperlabs/flow-go/issues/6398
- BN2 https://github.com/dapperlabs/flow-go/issues/6341
- TPS  https://github.com/dapperlabs/flow-go/issues/6296
- [Zero Quarantined Networking Layer Tests Epic](https://github.com/onflow/flow-go/issues/4816)

### **Consensus (Dynamic Protocol State) - Alex H**

**Done last sprint for Dynamic Protocol State**



**Done last sprint (other topics)**



**Next Sprint**
- Long list remaining tech todos for Dynamic Protocol State ([epic #4649](https://github.com/onflow/flow-go/issues/4649))
- continue automation of Cruise Control system


### **Infra - JP**

**Done last sprint**


**This Sprint**

************Node Hosting************
- Assist with Canary Sporks
- As needed, assist with HCU
- Create documentation for managing Dapper nodes
- Create automation for managing Dapper nodes

************FF Migration************
- Create strategy for observability migration
- Discover and configure tool for multi-destination logging
- Create strategy for GCP project migration to new GCP account

### Key Release Dates & Breaking Changes

- Mainnet/Testnet Spork dates 
  - Next spork - TBD (~Q1 2024)
- Second [Governance Working Group](https://github.com/onflow/gwg) meeting 9/28 (Tuesday)
