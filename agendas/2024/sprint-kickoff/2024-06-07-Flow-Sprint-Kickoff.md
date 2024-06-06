# Overview

### Team Wins 🎉

- 


### General updates

### OOO
- [Full List](https://www.notion.so/flowfoundation/de89aa4e79364216a665888335a1cdee?v=4de18b26f60d4bae8f62724dddcce260)

---

### Mainnet Uptime - Last 14 days (xx/yy/24 to xx/yy/24) \[Vishal]

|                         | Target | Current Score | Error budget used |
|:------------------------|:------:|:-------------:|:-----------------:|
| Collection Finalization | 99.9%   |    100%       |       0%         |
| Block Finalization      | 99.9%   |    100%       |       0%         |
| Transaction Execution   | 99.9%   |    99.95%     |       49.6%      |
| Block Sealing           | 99.9%   |    100%       |       0%         |
| Access API Liveness     | 99.9%   |    99.919%    |       81.0%      |

[SLO dashboards](https://flowfoundation.grafana.net/d/hgW3l-m4k/slo-dashboard?orgId=1&from=now-2w&to=now)

[YTD SLA: 100%](https://app.metrika.co/flow/dashboard/slas?tr=YTD)

## Incidents

## Core protocol incidents

### Mainnet

### Testnet

(Sev [definition](https://www.notion.so/dapperlabs/Incident-Priorities-Severity-Levels-b65d7682336c46baa025ee512fd3efa3))

### Key Release Dates & Breaking Changes
- Next Mainnet/Testnet network upgrade (spork):
  - Testnet: 20th June 2024
  - Mainnet: 29th July 2024
---

### FLIPs Tracker \[Kshitij]

|                         | Application | Cadence | Governance | Protocol | Total |  
|:------------------------|:------:|:-------------:|:-----------------:|:-----------------:|:-----------------:|
| Drafted     | 7 (+1)  |    8   |       0 (-1)         |       7          |        **22**        |
| Proposed    | 1  |    2     |       3 (+1)          |       0           |        **6** (+1)        |
| Accepted    | 2  |    1     |       1       |       1          |        **5**          |
| Rejected    | 0  |    0     |       1       |       0          |        **1**          |
| Implemented | 3 |    21    |       2       |       1           |        **7**          |
| Released    | 4  |    0     |       3       |       6          |        **13**          |
| Total       | **17** (+1)  |    **32**  |       **10** (+1)      |       **15**          |        **74** (+1)          |

**Updates**
* [Application] - FLIP 270: FCL `ServiceProvider` v1.1.0 (new)
* [Governance] - 5X Computation Limit Hike (moved from drafted to proposed)
---


# Working Group Updates

### **Cadence and Virtual Machine** \[Jan]
Cycle Objective(s):

1) Upgrade Mainnet to Crescendo Release with minimal impact on developers, to lower the barrier for cross chain liquidity on Flow
2) Calibrate Transaction fees so that they accurately reflect resource usage during execution and deploy as part of Crescendo to avoid future disruption.
3) Analyze execution runtime trends and risks to plan next set of scalability OKRs.

* Stretch-goals:
4) Expand testing capability of storehouse so that we can validate execution correctness and benchmark performance on Mainnet data
5) Design a new Trie to improve performance of update operation, reduce memory usage and size of proofs and to support more flexible proof queries.
6) Enable Concurrent Execution on one EN on Mainnet to validate correctness of the implementation (Detect execution forks)
7) Improve execution performance to mitigate the impact of adding metadata to token standards

**Done last sprint**

**This sprint**

DONE: Objective 1, KR 1: Enable Developers and the Flow Foundation to simulate Cadence 1.0 Contract upgrades
* All breaking changed released in a new CLI: v1.18.0-cadence-v1.0.0-preview.24 

IN PROGRESS: Objective 1, KR4: Testnet Upgrade to Crescendo Release
* Completed Testnet migration with both Atree inlining and Cadence 1.0.
* Completed [EVM Gateway development](https://github.com/onflow/flow-evm-gateway/issues/126) and [EVM Core development](https://github.com/onflow/flow-go/issues/5536) production readiness EPICs.
* Continue work on migration optimizations.

IN PROGRESS: Objective 2, KR 1: Update transaction fees weights for the execution operations on TN and MN
* Continue work on [Execution Effort Calibration](https://github.com/onflow/flow-go/issues/5598)

DONE: Objective 4, KR1: Execution node handles restart from spork root block reguardless of how many blocks it is behind
* Completed refactoring of Ingestion engine to [prevent EN entering crash loop](https://github.com/onflow/flow-go/issues/5298)

ON HOLD: Objective 3: Analyze execution runtime trends and risks to plan next set of scalability OKRs
* Continue work on making [Make TPS loader input more flexible](https://github.com/onflow/flow-go/issues/5490) for better analysis and tracking of performance data.

ON HOLD: * Start Atree optimization: [Adding support for lazy decoding of registers](https://github.com/onflow/atree/issues/341)

IN PROGRESS: EVM GW stability - investigating

ON HOLD: EVM GW Debug endpoint implementation

Cadence 1.0 Contract updates
- [Continue working through contract standards needing upgrade to C1.0](https://github.com/onflow/docs/issues/716)
- More reviews of bridge PRs and Cadence 1.0 changes 
- Writing additional tests for recent Cadence FLIP changes
- [Audit and update docs and tutorials for Cadence 1.0](https://github.com/onflow/docs/issues/531)

**On Hold**

Objective 2: Calibrate Transaction fees so that they accurately reflect resource usage during execution and deploy as part of Crescendo to avoid future disruption
- KR1: Update weights for the execution operations on TN and MN
  - [Execution effort recalibration - data collection](https://github.com/onflow/flow-go/issues/5026)


**Active Epics**

Objective 1: Upgrade Mainnet to Crescendo Release with minimal impact on developers, to lower the barrier for cross chain liquidity on Flow
- KR1: Enable Developers and the Flow Foundation to simulate Cadence 1.0 Contract upgrades
- KR4: Testnet Upgrade to Crescendo Release
- KR6: Develop & share with community a disaster recovery plan to address potential issues after migration to Crescendo Release.
- KR7: Decision on how to deal with contracts that have not been upgraded to Cadence 1.0 by developers.
Objective 3: Analyze execution runtime trends and risks to plan next set of scalability OKRs
- KR1: Measure execution state growth trends, identify future bottlenecks and prioritize by urgency

---

### **Core Protocol** \[Jerome]
Cycle Objective(s): 

* Translate crypto performance improvements to consensus block rate increase [DONE]
* Provide developers secure and non-rate limited way to access all of chain data (transactions, blocks, account balance, events, account balance etc) by locally running an access or an observer node
* Reduce CPU usage on Execution node by 30%
* Continue design and implementation of Sporkless Epoch Fallback Recovery solution [DONE]

**Done last Sprint:**

      
**This sprint**

* <ins>EFM Recovery</ins>
  - Finish https://github.com/onflow/flow-go/issues/5727
  - Onboard Jordan to EFM tasks
  - [Finish Epoch manager QC voting changes](https://github.com/onflow/flow-go/issues/5733) (implementing tests)
  - Ongoing review by SC team: [EpochRecover cadence transaction](https://github.com/onflow/flow-core-contracts/pull/420)
  - [Start Blocktime controller EFM changes])(https://github.com/onflow/flow-go/issues/5732)

* <ins>Data Availability:</ins>
  - Redeploy local index and script exec on QN bare metal instances
  - Start design work for DB pruning
  - Finish review of Event Streaming blog authored by KROK
  - KROK Team
    - Fix retries when using preferred-execution-nodes list ([Issue #5810](https://github.com/onflow/flow-go/issues/5810))
    - Add support for version beacon events to control script execution ([Issue #5757](https://github.com/onflow/flow-go/issues/5757))
    - Fix Access integration test ([Issue #5825](https://github.com/onflow/flow-go/issues/5825) - PR in review)
    - Add indexed height indicator in grpc metadata response ([Issue #4757](https://github.com/onflow/flow-go/issues/4757) - PR in review)

* <ins>Cryptography:</ins>
   - SPoCK aggregation: deeper look at the KOSK model
   - content piece on secure enclaves

* <ins>Rosetta:</ins>   
  - KROK [Finish Rosetta event hash check](https://github.com/onflow/rosetta/issues/41)
  - KROK [Continue Cadence Rosetta updates for Crescendo compatibility](https://github.com/onflow/rosetta/issues/41)
  - Review Flow EVM Rosetta implementation from contract developer

**On Hold**
* Deliver public roadmap & vision for technical protocol decentralization focusing on current challenges and upcoming updates for permissionless consensus on Flow.
* Implement BFT mitigations to enable 20 permissionless ANs

**Active Epics**

- Reinforcing Flow’s commitment to full protocol autonomy and scalability
- Improve network performance
- Improve network availability 
- Simplify community contributions to core protocol and maintainability
- Improve network reliability and data availability for dApp developers
- Data-driven Prioritization and Scaling Engineering

---

### **DeFi** \[Jerome]

Cycle Objective(s): 
- Resolving Circle's existing engineering improvements for USDC on Flow
- Cadence 1.0 DEX Prep - IncrementFi
- EVM partner onboarding: Pyth, DeBridge, Covalent, Etherscan
- Deliver Axelar bridge

**Done last sprint**
* PR to backfill of missing crypto algos to JVM SDK ready for review

**This sprint**
* Backfill missing API methods to JVM SDK

**On Hold**
- Axelar Cadence bridge release waiting confirmation of build partners, costs and timelines for EVM release, no engineering ongoing

**Active Epics**

- Establish Defi/Liquidity infrastructure for Cadence 1.0 update
- Ensure Flow has best-in-class on- and off-ramps for USDC liquidity across DeFi ecosystem
- Expand Flow DeFi ecosystem primitives and protocols

---

### **User Experience** \[Greg]

Cycle Objective(s):

- Bring Cadence 1.0 to market as part of the Crescendo release to minimize customer impact and developer effort
- Bring EVM on Flow to Market as part of the Crescendo release to increase liquidity and bring top-tier developer platforms to our network
- Use the Crescendo Release grow Flow's developer base and network activity

**Done last sprint**


**This sprint**
**Sprint goal focusing on: Flow Port Asset Bridge and Token Transfer, FCL x WalletConnect Improvements, Crescendo Rewards**

- [EPIC] Flow Port Asset Bridge [#284](https://github.com/onflow/flow-port/issues/284)
- [EPIC] EVM Docs [#654](https://github.com/onflow/docs/issues/654)
- [EPIC] Flow EVM Docs - Cadence Dev [#575](https://github.com/onflow/docs/issues/575)
- [EPIC]: Update Flow port for Cadence v1.0 instance [#279](https://github.com/onflow/flow-port/issues/279)
- [EPIC] FCL WalletConnect Improvements https://github.com/onflow/fcl-js/issues/1872

#### Smart Contract WG

- Reviews for Increment.fi
- Circle Support
- Add Solidity implementation using EVM randomness
- Discovery Design Inbox Standard

**On Hold**

---

### **Wallet** \[Jeff]

Cycle Objective(s): 

- Ensure there exists a wallet ecosystem supports FlowEVM
  - Release version 2.2 of Flow Wallet which supports FlowEVM
    - Support Authn / Authz / User Sign with Web3.js and WalletConnect
    - Support FT and NFT management cross VMs
    - FlowEVM onboarding and COA creation
  - Ensure commitments from key EVM wallet providers to support FlowEVM
    - Secure FlowEVM as an option in the network selector list for MetaMask.
    - Reach out to Coinbase wallet for a commitment to support FlowEVM
  - Ensure commitments from key EVM wallet providers to support FlowEVM
    - Reach out to Privy for a commitment to support FlowEVM
    - Reach out to Bastion for a commitment to support FlowEVM
    - Ensure awareness for existing Cadence aware wallet (aside from Flow Wallet) to support FlowEVM
  - Provide a design document outlining the steps existing Cadence aware wallets can take to support FlowEVM.
    - Reach out to Blocto for a commitment to support FlowEVM
    - Reach out to Shadow wallet for a commitment to support FlowEVM
    - Reach out to Magic for a commitment to support FlowEVM

- Promote safe, human readable transaction authorization on Flow
  - Secure a partnership with Blockaid to integrate their transaction simulation and security platform with FlowEVM.
  - Ensure the existing MetaMask Blockaid integration is compatible with FlowEVM.

- Modernize and improve FCL Discovery
  - Create a PRD and associated community bounty/grant for UI/UX improvements and analytics additions to FCL Discovery.

**Done last sprint**


**This sprint**
 
**On Hold**

- N/A

**Active Epics**

- TBD

---

### **Infra - JP**
Cycle Objective(s): 
- Solidify CloudFlare Migration plan and continue preparing for migration
- Continue assisting with DevEx migration to Cloud Run

**Done last sprint**

**Goal of this Sprint is to continue migration efforts** 
**This Sprint**
- Continue removing dependencies on CloudFlare & assist with account creation
- Continue assisting with DevEx migration to Cloud Run
- Improve synthetic alerting 
- Assist with migration and spork prep efforts

---

### **Governance and Tokenomics** \[Kshitij]
Cycle Objective(s): Transaction fees on EVM, increasing transaction fees and inflation reduction plan.

**Done last sprint**

**This sprint**
- Set up 3 FF nodes via Figment
- Onboard Doodles and Cryptotoys as SN operators
- Kick off Flowty's consensus node
- Assess decentralization situation and the way forward
- Review, finalize and take a decision on computation limit FLIP 
- Transaction fee partner impact discussion and way forward with FF leadership
- R&D and planning with Dete on post Crescendo surge pricing
- Continue working on node operator branding and logos with .find team
  
**On Hold**


**Active Epics**

- N/A