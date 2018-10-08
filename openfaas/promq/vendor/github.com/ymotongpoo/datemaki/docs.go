// Package datemaki is a flexible and fancy datetime parsing library.
//
// This is work in progress. This package aims to handle datetime
// expressed in more natural languages and convert them into
// time.Time instances.
//
// Currently this package supports following datetime expression, mainly
// used in Git commands.
//
// * Past datetime described in ago
//   * 2 seconds ago
//   * 3 minutes ago
//   * 4 hours ago
//   * 5 days ago
//   * 1 week ago
//   * 2 months ago
//   * 1 year, 3 months ago
//   * 1.year.4.months.ago
//   * 2.years.ago
//
// * Relative date
//   * now
//   * today
//   * yesterday
//   * last friday
//
// * Relative date and fixed time
//   * noon yesterday
//   * tea yesterday
//   * midnight today
//   * 3pm today
//   * 2am last friday
//   * 19:00 yesterday (under implementation)
//   * 10am
//
// * Absolute datetime
//   * August 6th
//   * 06/05/2009
//   * 06.05.2009
//   * Feb 28, 4AM
//   * 2AM Jun 4
//   * 6AM, June 7, 2009
//   * 2008-12-01
package datemaki
