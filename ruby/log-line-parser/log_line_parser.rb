class LogLineParser
  attr_reader :log_level, :message

  def initialize(line)
    m = /\[(?<log_level>[A-Z]+)\]:\s*(?<message>[\w| ]+\b)/.match(line)

    @log_level = m[:log_level].downcase
    @message = m[:message]
  end

  def reformat
    "#{@message} (#{@log_level})"
  end
end
