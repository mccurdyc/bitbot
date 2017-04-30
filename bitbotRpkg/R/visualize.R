
#' FUNCTION: visualize_algorithm_worth
#'
#' Bar chart with the worth of each algorithm at a given time stamp for a particular configuration.
#' @export

visualize_algorithm_worth <- function(d) {
  p1 <- d %>% visualize_plot_worth()
  name1 <- "../graphics/from-data/algorithms_worth_high.pdf"
  visualize_save_graphic(name1, p1, 8, 8)

  p2 <- d %>% visualize_plot_worth_pres()
  name2 <- "../graphics/from-data/algorithms_worth_high_pres.pdf"
  visualize_save_graphic(name2, p2, 12, 6)
}


#' FUNCTION: visualize_plot_worth
#'
#' @export

visualize_plot_worth <- function(d) {
  p <- ggplot2::ggplot(d, ggplot2::aes(x = date, y = worth, group=interaction(date, algorithm))) +
    ggplot2::theme(strip.background = element_blank(), panel.border = element_rect(colour = "black")) +
    ggplot2::geom_boxplot(ggplot2::aes(fill = algorithm)) +
    ggplot2::facet_wrap(small ~ large, labeller = ggplot2::label_both) +
    ggplot2::theme_bw(base_size = 10) +
    ggplot2::theme(axis.text.x = ggplot2::element_text(angle = 45, hjust = 1, size = 10)) +
    ggplot2::theme(axis.text.y = ggplot2::element_text(angle = 45, hjust = 1, size = 10)) +
    ggplot2::xlab("Date") +
    ggplot2::ylab("Worth")
  return(p)
}

#' FUNCTION: visualize_plot_worth_pres
#'
#' @export

visualize_plot_worth_pres <- function(d) {
  p <- ggplot2::ggplot(d, ggplot2::aes(x = date, y = worth, group = interaction(date, algorithm))) +
    ggplot2::geom_boxplot(ggplot2::aes(colour = algorithm)) +
    ggplot2::theme_bw(base_size = 8) +
    ggplot2::theme(strip.background = element_blank(), panel.border = element_rect(colour = "black"), legend.position="none") +
    ggplot2::theme(axis.text.x = ggplot2::element_text(angle = 45, hjust = 1, size = 15)) +
    ggplot2::theme(axis.text.y = ggplot2::element_text(angle = 45, hjust = 1, size = 15)) +
    ggplot2::theme(axis.title.x = ggplot2::element_text(size = 25)) +
    ggplot2::theme(axis.title.y = ggplot2::element_text(size = 25)) +
    ggplot2::theme(panel.background = element_rect(fill = "transparent", colour = NA)) +
    ggplot2::theme(plot.background = element_rect(fill = "transparent", colour = NA)) +
    ggplot2::xlab("Date") +
    ggplot2::ylab("Worth")
  return(p)
}

#' FUNCTION: visualize_save_graphic
#'
#' Saves the provided graphic to the provided name.
#' @export

visualize_save_graphic <- function(save_name, save_plot, w, h) {
  ggplot2::ggsave(save_name, save_plot, width = w, height = h)
}
