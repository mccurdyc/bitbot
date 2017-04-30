#' FUNCTION: summarize_algorithm_worth
#'
#' A function for determining which algorithm is better on average at generating more 'worth' in one month's time.
#' @export

summarize_algorithm_worth <- function(d) {
  d <- d %>% collect_funds_small_large()
  dt <- d %>% dplyr::summarize(greedy_mean_worth = mean(a1_worth), random_mean_worth = mean(a2_worth))
  return(dt)
}

#' FUNCTION: collect_funds_small_large
#'
#' Group data by initial funds, the small and large date ranges for calculating simple moving averages.
#' @export

collect_funds_small_large <- function(d) {
  dt <- d %>% dplyr::group_by(initial_funds, small, large) %>%
  return(dt)
}

#' FUNCTION: display_algorithm_worth
#'
#' A function for displaying the worth of each algorithm for many configuration scenarios.
#' @export

display_algorithm_worth <- function(d) {
  d <- d %>% collect_funds_small_large()
  d %>% visualize_algorithm_worth()
}


