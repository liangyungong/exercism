class LogLineParser

  LEVELS = %w{INFO WARNING ERROR}

  def initialize(line)
    @line = line.strip
  end

  def message
    l = @line.clone
    LEVELS.each do |level|
      l.slice! "[#{level}]: "
    end
    l.strip
  end

  def log_level
    l = @line.clone
    LEVELS.each do |level|
      if l.include? "[#{level}]: "
        return level.downcase
      end
    end
  end

  def reformat
    l = @line.clone
    LEVELS.each do |level|
      if l.include? "[#{level}]: "
        l.slice! "[#{level}]: "
        return "#{l.strip} (#{level.downcase})"
      end
    end
  end
end
