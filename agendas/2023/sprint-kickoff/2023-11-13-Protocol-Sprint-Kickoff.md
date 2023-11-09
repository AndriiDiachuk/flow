# Team Wins 🎉

- safe randomness tools are deployed and working on Mainnet - launch page, documentation and first tutorial are live 🎲
- Script execution working on devnet! Will enable on mainnet this sprint.
- Event streaming available for REST on mainnet. (working with 4d on fcl integration)
- Lots of AccessAPI quality of life improvements deployed to mainnet. (grpc compression, optional CCF, historic AN tx result cache)

### Mainnet Uptime SLO - Last 14 days (10/13 to 10/27)

|                         | Target | Current Score | Error budget used |
|:------------------------|:------:|:-------------:|:-----------------:|
| Collection Finalization | 99.9%  |    100%       |       0%          |
| Block Finalization      | 99.9%  |    99.896%    |       104%        |
| Transaction Execution   | 99.9%  |    99.906%    |       94.3%       |
| Block Sealing           | 99.9%  |    99.816%    |       174%        |
| Access API Liveness     | 99.9%  |    99.722%    |       278%        |

#### Incidents



### **Performance Pod Sprint Objective - Jan B**

**Done last sprint**


**This sprint**

- [Continue Cadence integration to use Atree register inlining](https://github.com/onflow/cadence/issues/2809)
- Continue testing [Atree register inlining migration](https://github.com/onflow/flow-go/pull/4633)
- Continue implementation of [Storehouse first milestone](https://github.com/onflow/flow-go/issues/4682) (execution state on disk)

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

**This sprint**

- continue support EVM on FLow initiative.
- Continuing with Stable Cadence scope / discussions
    - Ongoing FLIPs:
        - Last FLIP to be opened - Update on entitlements on Attachments
- Continue work on Cadence 1.0 migrations.
- Continue Stable Cadence Docs update and knocking tasks off the [tech debt list](https://github.com/onflow/cadence/issues/2642)
 
**On Hold**
- Discussion of the re-entrancy edge cases

**Active Epics**
- [Stable Cadence (aka Cadence 1.0)](https://github.com/onflow/cadence/issues/2642)


### **Access & Data Availability - Peter A**
Objective: Make execution data and script execution available on Edge nodes.

**Done last sprint**

Script Execution:

- [Ledger] Add special handling for global register keys - [PR 4942](https://github.com/onflow/flow-go/pull/4942)
- [Execution] Return OutOfRange instead of Internal when account block is not cached - [PR 4917](https://github.com/onflow/flow-go/pull/4917)
- [Access] Allow get blocks script calls - [PR 4894](https://github.com/onflow/flow-go/pull/4894)
- [Access] Script execution coded errors - [PR 4895](https://github.com/onflow/flow-go/pull/4895)
- [Access] Get account bugfix with tests - [PR 4862](https://github.com/onflow/flow-go/pull/4862)
- [Access] Validate addresses match network in rest api - [PR 4930](https://github.com/onflow/flow-go/pull/4930)
- [Access] Add metrics for script exec failure from missing data - [PR 4907](https://github.com/onflow/flow-go/pull/4907)
- [Access] Improve logging and validation in local script exec - [PR 4920](https://github.com/onflow/flow-go/pull/4920)
- [Access] Improve script exec compare logging - [PR 4936](https://github.com/onflow/flow-go/pull/4936)
- [Access] Cleanup script execution comparisons - [PR 4956](https://github.com/onflow/flow-go/pull/4956)

Access API:

- [Access] Allow all origins by default on websockets connections - [PR 4954](https://github.com/onflow/flow-go/pull/4954)
- [Flow-Go-SDK] Use CCF encoding when requesting events from AccessAPI - [PR 501](https://github.com/onflow/flow-go-sdk/pull/501)

Misc:

- [Access] Add wait in integration tests for index to be synced - [PR 4902](https://github.com/onflow/flow-go/pull/4902)
- [Collection] Make QC Voter more resiliant to access node instability - [PR 4924](https://github.com/onflow/flow-go/pull/4924)

**This sprint**

- Script Execution on ANs
  - Add GetRegisters API endpoint to ExecutionData API - [Issue 4756](https://github.com/onflow/flow-go/issues/4756)
  - Analyze performance issues observed on devnet - [Issue 4953](https://github.com/onflow/flow-go/issues/4953)
  - Deploy to mainnet (in comparison mode) and continue analyzing results and performance issues as they come up.
- Misc
  - Work with 4d on getting event streaming and CCF into libraries.
  - Validate new features (historic result cache, compression, etc)

**Active Epics**

- Script Execution on Access Node - [Issue 4637](https://github.com/onflow/flow-go/issues/4637)
- Integrate local execution state indexes into Access API - [Issue 4750](https://github.com/onflow/flow-go/issues/4750)


### **Permissionless Network - Yahya H**

**Done last sprint**
- [Investigating and fixing goroutine leakage on `mainnet23`](https://github.com/dapperlabs/flow-go/issues/6871) [PR4846](https://github.com/onflow/flow-go/pull/4846)
- [Addressing technical debts for sync engine ALSP integration](https://github.com/onflow/flow-go/pull/4842) 

**Ongoing (last and next sprint)**
- [Investigated and fixed AN-LN streaming issue on `mainnet23`](https://github.com/dapperlabs/flow-go/issues/6895) [PR4875](https://github.com/onflow/flow-go/pull/4875)
- [Optimizing memory-intensive RPC inspection operations](https://github.com/dapperlabs/flow-go/issues/6870)
- [Gossip scoring to support additional cluster prefixed control messages](https://github.com/dapperlabs/flow-internal/issues/1889) [PR4857](https://github.com/onflow/flow-go/pull/4857)
- [Implement Specific Decay per Peer ID in GossipSubSpamRecord for Improved Spam Mitigation](https://github.com/dapperlabs/flow-go/issues/6662)

**Next Sprint**
- [Balancing the inbound and outbound resource limits with backpressure](https://github.com/dapperlabs/flow-go/issues/6896)
- [Apply Penalty to Misbehaving Peers Based on Count and Err in InvCtrlMsgNotif](https://github.com/dapperlabs/flow-go/issues/6664)
- [Increase test coverage for BFTune ingress unit tests](https://github.com/dapperlabs/flow-go/issues/6883)
- [[BFT Testing] Refactor Orchestrator lock contension to use worker pools](https://github.com/dapperlabs/flow-go/issues/6884)
- [[CI][Testing] Increase GitHub CI runners for resource intensive tests](https://github.com/dapperlabs/flow-go/issues/6894)

**Active Epics**

- https://github.com/dapperlabs/flow-go/issues/6287
- https://github.com/dapperlabs/flow-go/issues/6468
- BFT https://github.com/dapperlabs/flow-go/issues/6142
- BFT https://github.com/dapperlabs/flow-go/issues/6398
- BN2 https://github.com/dapperlabs/flow-go/issues/6341
- TPS  https://github.com/dapperlabs/flow-go/issues/6296
- [Zero Quarantined Networking Layer Tests Epic](https://github.com/onflow/flow-go/issues/4816)

### **Consensus (Dynamic Protocol State) - Alex H**

**Done last sprint**

**Ongoing** (last & next sprint)

- Consolidating exploratory research documents to reflect final design (requires change to reflect removal of dynamic weight)
- Refactoring and refining the interfaces and data structures for clarity (👉 [epic #4649](https://github.com/onflow/flow-go/issues/4649))
- Handle high priority [TODOs](https://github.com/onflow/flow-go/issues/4649), specifically: change how protocol state behaves in case of invalid state transition and replace dynamic weight with participation status.


### **Infra - JP**

**Done last sprint**

**This Sprint**
- Update Ansible automation for Dapper nodes
- Prepare monitoring & alerting for Dapper Nodes
- Prepare Dapper infra/keys for Mainnet spork
- Create Flow Foundation infrastructure & configuration for Mainnet spork
- Create Dapper infrastructure & configuration for Mainnet spork

************Node Hosting************
### Key Release Dates & Breaking Changes

- Mainnet/Testnet Spork dates 
  - Next spork