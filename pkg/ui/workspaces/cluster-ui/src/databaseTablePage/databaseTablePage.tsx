// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

import React from "react";
import { Col, Row, Tabs } from "antd";
import "antd/lib/col/style";
import "antd/lib/row/style";
import "antd/lib/tabs/style";
import { RouteComponentProps } from "react-router-dom";
import classNames from "classnames/bind";
import { Tooltip } from "antd";
import "antd/lib/tooltip/style";
import { Heading } from "@cockroachlabs/ui-components";

import { Anchor } from "src/anchor";
import { Breadcrumbs } from "src/breadcrumbs";
import { CaretRight } from "src/icon/caretRight";
import { StackIcon } from "src/icon/stackIcon";
import { SqlBox } from "src/sql";
import { ColumnDescriptor, SortSetting, SortedTable } from "src/sortedtable";
import {
  SummaryCard,
  SummaryCardItem,
  SummaryCardItemBoolSetting,
} from "src/summaryCard";
import * as format from "src/util/format";
import { syncHistory, tableStatsClusterSetting } from "src/util";

import styles from "./databaseTablePage.module.scss";
import { commonStyles } from "src/common";
import { baseHeadingClasses } from "src/transactionsPage/transactionsPageClasses";
import moment, { Moment } from "moment";
import { Search as IndexIcon } from "@cockroachlabs/icons";
import { Link } from "react-router-dom";
import classnames from "classnames/bind";
import booleanSettingStyles from "../settings/booleanSetting.module.scss";
import { CircleFilled } from "../icon";
import { performanceTuningRecipes } from "src/util/docs";
import { DATE_FORMAT_24_UTC, DATE_FORMAT } from "src/util/format";
const cx = classNames.bind(styles);
const booleanSettingCx = classnames.bind(booleanSettingStyles);

const { TabPane } = Tabs;

// We break out separate interfaces for some of the nested objects in our data
// so that we can make (typed) test assertions on narrower slices of the data.
//
// The loading and loaded flags help us know when to dispatch the appropriate
// refresh actions.
//
// The overall structure is:
//
//   interface DatabaseTablePageData {
//     databaseName: string;
//     name: string;
//     details: { // DatabaseTablePageDataDetails
//       loading: boolean;
//       loaded: boolean;
//       createStatement: string;
//       replicaCount: number;
//       indexNames: string[];
//       grants: {
//         user: string;
//         privilege: string;
//       }[];
//     };
//     stats: { // DatabaseTablePageDataStats
//       loading: boolean;
//       loaded: boolean;
//       sizeInBytes: number;
//       rangeCount: number;
//       nodesByRegionString: string;
//     };
//     indexStats: { // DatabaseTablePageIndexStats
//       loading: boolean;
//       loaded: boolean;
//       stats: {
//         indexName: string;
//         totalReads: number;
//         lastUsed: Moment;
//         lastUsedType: string;
//       }[];
//       lastReset: Moment;
//     };
//   }
export interface DatabaseTablePageData {
  databaseName: string;
  name: string;
  details: DatabaseTablePageDataDetails;
  stats: DatabaseTablePageDataStats;
  indexStats: DatabaseTablePageIndexStats;
  showNodeRegionsSection?: boolean;
  automaticStatsCollectionEnabled?: boolean;
}

export interface DatabaseTablePageDataDetails {
  loading: boolean;
  loaded: boolean;
  createStatement: string;
  replicaCount: number;
  indexNames: string[];
  grants: Grant[];
  statsLastUpdated?: Moment;
  totalBytes: number;
  liveBytes: number;
  livePercentage: number;
}

export interface DatabaseTablePageIndexStats {
  loading: boolean;
  loaded: boolean;
  stats: IndexStat[];
  lastReset: Moment;
}

interface IndexStat {
  indexName: string;
  totalReads: number;
  lastUsed: Moment;
  lastUsedType: string;
  indexRecommendations: IndexRecommendation[];
}

interface IndexRecommendation {
  type: string;
  reason: string;
}

interface Grant {
  user: string;
  privilege: string;
}

export interface DatabaseTablePageDataStats {
  loading: boolean;
  loaded: boolean;
  sizeInBytes: number;
  rangeCount: number;
  nodesByRegionString?: string;
}

export interface DatabaseTablePageActions {
  refreshTableDetails: (database: string, table: string) => void;
  refreshTableStats: (database: string, table: string) => void;
  refreshSettings: () => void;
  refreshIndexStats?: (database: string, table: string) => void;
  resetIndexUsageStats?: (database: string, table: string) => void;
  refreshNodes?: () => void;
}

export type DatabaseTablePageProps = DatabaseTablePageData &
  DatabaseTablePageActions &
  RouteComponentProps;

interface DatabaseTablePageState {
  sortSetting: SortSetting;
  tab: string;
}

class DatabaseTableGrantsTable extends SortedTable<Grant> {}
class IndexUsageStatsTable extends SortedTable<IndexStat> {}

export class DatabaseTablePage extends React.Component<
  DatabaseTablePageProps,
  DatabaseTablePageState
> {
  constructor(props: DatabaseTablePageProps) {
    super(props);

    const { history } = this.props;
    const searchParams = new URLSearchParams(history.location.search);
    const defaultTab = searchParams.get("tab") || "overview";

    this.state = {
      sortSetting: {
        ascending: true,
      },
      tab: defaultTab,
    };
  }

  onTabChange = (tab: string): void => {
    this.setState({ tab });
    syncHistory(
      {
        tab: tab,
      },
      this.props.history,
    );
  };

  componentDidMount(): void {
    this.refresh();
  }

  componentDidUpdate(): void {
    this.refresh();
  }

  private refresh() {
    if (this.props.refreshNodes != null) {
      this.props.refreshNodes();
    }
    if (!this.props.details.loaded && !this.props.details.loading) {
      return this.props.refreshTableDetails(
        this.props.databaseName,
        this.props.name,
      );
    }

    if (!this.props.stats.loaded && !this.props.stats.loading) {
      return this.props.refreshTableStats(
        this.props.databaseName,
        this.props.name,
      );
    }

    if (!this.props.indexStats.loaded && !this.props.indexStats.loading) {
      return this.props.refreshIndexStats(
        this.props.databaseName,
        this.props.name,
      );
    }

    if (this.props.refreshSettings != null) {
      this.props.refreshSettings();
    }
  }

  minDate = moment.utc("0001-01-01"); // minimum value as per UTC

  private changeSortSetting(sortSetting: SortSetting) {
    this.setState({ sortSetting });
  }

  private getLastResetString() {
    const lastReset = this.props.indexStats.lastReset;
    if (lastReset.isSame(this.minDate)) {
      return "Last reset: Never";
    } else {
      return "Last reset: " + lastReset.format(DATE_FORMAT_24_UTC);
    }
  }

  private getLastUsedString(indexStat: IndexStat) {
    // This case only occurs when we have no reads, resets, or creation time on
    // the index.
    if (indexStat.lastUsed.isSame(this.minDate)) {
      return "Never";
    }
    return `Last ${indexStat.lastUsedType}: ${indexStat.lastUsed.format(
      DATE_FORMAT,
    )}`;
  }

  private renderIndexRecommendations = (
    indexStat: IndexStat,
  ): React.ReactNode => {
    const classname =
      indexStat.indexRecommendations.length > 0
        ? "index-recommendations-icon__exist"
        : "index-recommendations-icon__none";

    if (indexStat.indexRecommendations.length === 0) {
      return (
        <div>
          <CircleFilled className={cx(classname)} />
          <span>None</span>
        </div>
      );
    }
    return indexStat.indexRecommendations.map(recommendation => {
      let text: string;
      switch (recommendation.type) {
        case "DROP_UNUSED":
          text = "Drop unused index";
      }
      // TODO(thomas): using recommendation.type as the key seems not good.
      //  - if it is possible for an index to have multiple recommendations of the same type
      //  this could cause issues.
      return (
        <Tooltip
          key={recommendation.type}
          placement="bottom"
          title={
            <div className={cx("index-recommendations-text__tooltip-anchor")}>
              {recommendation.reason}{" "}
              <Anchor href={performanceTuningRecipes} target="_blank">
                Learn more
              </Anchor>
            </div>
          }
        >
          <CircleFilled className={cx(classname)} />
          <span className={cx("index-recommendations-text__border")}>
            {text}
          </span>
        </Tooltip>
      );
    });
  };

  private indexStatsColumns: ColumnDescriptor<IndexStat>[] = [
    {
      name: "indexes",
      title: "Indexes",
      hideTitleUnderline: true,
      className: cx("index-stats-table__col-indexes"),
      cell: indexStat => (
        <Link
          to={`${this.props.name}/index/${indexStat.indexName}`}
          className={cx("icon__container")}
        >
          <IndexIcon className={cx("icon--s", "icon--primary")} />
          {indexStat.indexName}
        </Link>
      ),
      sort: indexStat => indexStat.indexName,
    },
    {
      name: "total reads",
      title: "Total Reads",
      hideTitleUnderline: true,
      cell: indexStat => format.Count(indexStat.totalReads),
      sort: indexStat => indexStat.totalReads,
    },
    {
      name: "last used",
      title: "Last Used (UTC)",
      hideTitleUnderline: true,
      className: cx("index-stats-table__col-last-used"),
      cell: indexStat => this.getLastUsedString(indexStat),
      sort: indexStat => indexStat.lastUsed,
    },
    {
      name: "index recommendations",
      title: (
        <Tooltip
          placement="bottom"
          title="Index recommendations will appear if the system detects improper index usage, such as the occurrence of unused indexes. Following index recommendations may help improve query performance."
        >
          Index recommendations
        </Tooltip>
      ),
      cell: this.renderIndexRecommendations,
      sort: indexStat => indexStat.indexRecommendations.length,
    },
  ];

  private grantsColumns: ColumnDescriptor<Grant>[] = [
    {
      name: "username",
      title: (
        <Tooltip placement="bottom" title="The user name.">
          User Name
        </Tooltip>
      ),
      cell: grant => grant.user,
      sort: grant => grant.user,
    },
    {
      name: "privilege",
      title: (
        <Tooltip placement="bottom" title="The list of grants for the user.">
          Grants
        </Tooltip>
      ),
      cell: grant => grant.privilege,
      sort: grant => grant.privilege,
    },
  ];

  formatMVCCInfo = (
    details: DatabaseTablePageDataDetails,
  ): React.ReactElement => {
    return (
      <>
        {format.Percentage(details.livePercentage, 1, 1)}
        {" ("}
        <span className={cx("bold")}>
          {format.Bytes(details.liveBytes)}
        </span>{" "}
        live data /{" "}
        <span className={cx("bold")}>{format.Bytes(details.totalBytes)}</span>
        {" total)"}
      </>
    );
  };

  render(): React.ReactElement {
    return (
      <div className="root table-area">
        <section className={baseHeadingClasses.wrapper}>
          <Breadcrumbs
            items={[
              { link: "/databases", name: "Databases" },
              { link: `/database/${this.props.databaseName}`, name: "Tables" },
              {
                link: `/database/${this.props.databaseName}/table/${this.props.name}`,
                name: `Table: ${this.props.name}`,
              },
            ]}
            divider={
              <CaretRight className={cx("icon--xxs", "icon--primary")} />
            }
          />

          <h3
            className={`${baseHeadingClasses.tableName} ${cx(
              "icon__container",
            )}`}
          >
            <StackIcon className={cx("icon--md", "icon--title")} />
            {this.props.name}
          </h3>
        </section>

        <section className={(baseHeadingClasses.wrapper, cx("tab-area"))}>
          <Tabs
            className={commonStyles("cockroach--tabs")}
            onChange={this.onTabChange}
            activeKey={this.state.tab}
          >
            <TabPane tab="Overview" key="overview" className={cx("tab-pane")}>
              <Row gutter={18}>
                <Col className="gutter-row" span={18}>
                  <SqlBox value={this.props.details.createStatement} />
                </Col>
              </Row>

              <Row gutter={18}>
                <Col className="gutter-row" span={8}>
                  <SummaryCard className={cx("summary-card")}>
                    <SummaryCardItem
                      label="Size"
                      value={format.Bytes(this.props.stats.sizeInBytes)}
                    />
                    <SummaryCardItem
                      label="Replicas"
                      value={this.props.details.replicaCount}
                    />
                    <SummaryCardItem
                      label="Ranges"
                      value={this.props.stats.rangeCount}
                    />
                    <SummaryCardItem
                      label="% of Live Data"
                      value={this.formatMVCCInfo(this.props.details)}
                    />
                    {this.props.details.statsLastUpdated && (
                      <SummaryCardItem
                        label="Table Stats Last Updated"
                        value={this.props.details.statsLastUpdated.format(
                          DATE_FORMAT_24_UTC,
                        )}
                      />
                    )}
                    {this.props.automaticStatsCollectionEnabled != null && (
                      <SummaryCardItemBoolSetting
                        label="Auto Stats Collection"
                        value={this.props.automaticStatsCollectionEnabled}
                        toolTipText={
                          <span>
                            {" "}
                            Automatic statistics can help improve query
                            performance. Learn how to{" "}
                            <Anchor
                              href={tableStatsClusterSetting}
                              target="_blank"
                              className={booleanSettingCx(
                                "crl-hover-text__link-text",
                              )}
                            >
                              manage statistics collection
                            </Anchor>
                            .
                          </span>
                        }
                      />
                    )}
                  </SummaryCard>
                </Col>

                <Col className="gutter-row" span={10}>
                  <SummaryCard className={cx("summary-card")}>
                    {this.props.showNodeRegionsSection && (
                      <SummaryCardItem
                        label="Regions/nodes"
                        value={this.props.stats.nodesByRegionString}
                      />
                    )}
                    <SummaryCardItem
                      label="Database"
                      value={this.props.databaseName}
                    />
                    <SummaryCardItem
                      label="Indexes"
                      value={this.props.details.indexNames.join(", ")}
                      className={cx("database-table-page__indexes--value")}
                    />
                  </SummaryCard>
                </Col>
              </Row>
              <Row gutter={18}>
                <SummaryCard
                  className={cx("summary-card", "index-stats__summary-card")}
                >
                  <div className={cx("index-stats__header")}>
                    <Heading type="h5">Index Stats</Heading>
                    <div className={cx("index-stats__reset-info")}>
                      <Tooltip
                        placement="bottom"
                        title="Index stats accumulate from the time the index was created or had its stats reset. Clicking ‘Reset all index stats’ will reset index stats for the entire cluster."
                      >
                        <div
                          className={cx("index-stats__last-reset", "underline")}
                        >
                          {this.getLastResetString()}
                        </div>
                      </Tooltip>
                      <div>
                        <a
                          className={cx(
                            "action",
                            "separator",
                            "index-stats__reset-btn",
                          )}
                          onClick={() =>
                            this.props.resetIndexUsageStats(
                              this.props.databaseName,
                              this.props.name,
                            )
                          }
                        >
                          Reset all index stats
                        </a>
                      </div>
                    </div>
                  </div>
                  <IndexUsageStatsTable
                    className="index-stats-table"
                    data={this.props.indexStats.stats}
                    columns={this.indexStatsColumns}
                    sortSetting={this.state.sortSetting}
                    onChangeSortSetting={this.changeSortSetting.bind(this)}
                    loading={this.props.indexStats.loading}
                  />
                </SummaryCard>
              </Row>
            </TabPane>
            <TabPane tab="Grants" key="grants" className={cx("tab-pane")}>
              <DatabaseTableGrantsTable
                data={this.props.details.grants}
                columns={this.grantsColumns}
                sortSetting={this.state.sortSetting}
                onChangeSortSetting={this.changeSortSetting.bind(this)}
                loading={this.props.details.loading}
              />
            </TabPane>
          </Tabs>
        </section>
      </div>
    );
  }
}
